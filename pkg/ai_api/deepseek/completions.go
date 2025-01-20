package deepseek

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
)

type CompletionsReq openai.ChatCompletionRequest

func NewCompletionReq(ques string) CompletionsReq {
	req := CompletionsReq{
		Model: DefaultModel,
	}
	req.AddUserMsg(ques)
	return req
}

func (r *CompletionsReq) Valid() error {
	if len(r.Messages) == 0 {
		return ErrNoMessage
	}
	if len(r.Model) == 0 {
		r.Model = DefaultModel
	}
	if !util.InRange(r.FrequencyPenalty, -2.0, 2.0) {
		return ErrFrequencyPenalty
	}

	return nil
}

func (r *CompletionsReq) AddMsgs(msg ...openai.ChatCompletionMessage) *CompletionsReq {
	r.Messages = append(r.Messages, msg...)
	return r
}

func (r *CompletionsReq) AddUserMsg(ques string) *CompletionsReq {
	r.Messages = append(r.Messages, openai.ChatCompletionMessage{
		Role:    RoleUser,
		Content: ques,
	})
	return r
}

const DefaultModel = "deepseek-chat"
const (
	RoleSys  = "system"
	RoleUser = "user"
	RoleAssi = "assistant"
	RoleTool = "tool"
)

type CompletionsResp openai.ChatCompletionResponse

func (c *Client) Completions(topic string, req CompletionsReq) (resp CompletionsResp, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var history []openai.ChatCompletionMessage
	err = c.store.View(func(tx *bolt.Tx) error {
		topicHistory := tx.Bucket(_BucketCompletions).Get(S2b(topic))
		if len(topicHistory) == 0 {
			return nil
		}
		err = Dec(topicHistory, &history)
		if err != nil {
			return err
		}
		req.AddMsgs(history...)
		return nil
	})
	if err != nil {
		return resp, errors.Wrap(err, "lookup local history")
	}

	r := c.cli.R()
	rresp, err := r.SetBody(req).
		Post("/chat/completions")
	if err != nil {
		return
	}
	if rresp.StatusCode() != http.StatusOK {
		errH := util.NewErrHandle()
		errH.Has(ErrRespNotOK)
		errH.HasStr("caused by ->")
		errH.Has(err)
		return resp, errH.All(" ")
	}
	err = Dec(rresp.Body(), &resp)
	if err != nil {
		return resp, errors.Wrap(err, "request to deepseek api")
	}

	err = c.store.Update(func(tx *bolt.Tx) error {
		totalMsg, err := Enc(append(append(history, req.Messages...),
			lo.Map(resp.Choices, func(item openai.ChatCompletionChoice, _ int) openai.ChatCompletionMessage { return item.Message })...))
		if err != nil {
			return err
		}
		err = tx.Bucket(_BucketCompletions).Put(S2b(topic), totalMsg)
		return err
	})
	if err != nil {
		return resp, errors.Wrap(err, "store latest history")
	}

	return
}

var _BucketCompletions = []byte("bkt_completions")
