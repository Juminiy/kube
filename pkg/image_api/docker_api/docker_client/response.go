package docker_client

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func EncE(v any) ([]byte, error) {
	return safe_json.Goccy().Marshal(v)
}
func DecE(b []byte, v any) error {
	return safe_json.Goccy().Unmarshal(b, v)
}

type EventResp struct {
	Status int

	// Status is 200, Message is not null
	Message       []*jsonmessage.JSONMessage
	MessageDecErr []string // unmarshal Message error

	// Status is not 200, ErrMessage is not null
	ErrMessage       []*jsonmessage.JSONError
	ErrMessageDecErr []string // unmarshal ErrMessage error
}

func (r *EventResp) WrapParse(resp *resty.Response, err error) (EventResp, error) {
	if err != nil { // http Request
		return *r, err // Client Error
	}
	parsed := r.Parse(resp) // http Response
	if parsed.err() {
		return *parsed, r.Err() // Server Response Error
	}
	return *parsed, nil
}

func (r *EventResp) Parse(resp *resty.Response) *EventResp {
	r.Status = resp.StatusCode()
	return r.parse(resp.Body())
}

func (r *EventResp) ParseBytes(bs []byte) *EventResp { // Has Done Error and Status Code
	return r.parse(bs)
}

func (r *EventResp) parse(bs []byte) *EventResp {
	if len(bs) == 0 {
		return r
	}
	sbs := bytes.Split(bs, []byte{'\r', '\n'})
	if r.Status == http.StatusOK {
		r.Message = make([]*jsonmessage.JSONMessage, len(sbs))
		r.MessageDecErr = make([]string, len(sbs))
		for i := range sbs {
			if len(sbs[i]) == 0 {
				continue
			}
			msg := jsonmessage.JSONMessage{}
			if err := DecE(sbs[i], &msg); err == nil {
				r.Message[i] = &msg
			} else {
				r.MessageDecErr[i] = err.Error()
			}
		}
	} else {
		r.ErrMessage = make([]*jsonmessage.JSONError, len(sbs))
		r.ErrMessageDecErr = make([]string, len(sbs))
		for i := range sbs {
			if len(sbs[i]) == 0 {
				continue
			}
			errMsg := jsonmessage.JSONError{}
			if err := DecE(sbs[i], &errMsg); err == nil {
				r.ErrMessage[i] = &errMsg
			} else {
				r.ErrMessageDecErr[i] = err.Error()
			}
		}
	}

	return r
}

func (r *EventResp) Error() string {
	errH := util.NewErrHandle()
	errH.Has(fmt.Errorf("response status code: %d", r.Status))
	for _, msg := range r.Message {
		if msg == nil || msg.Error == nil {
			continue
		}
		errH.Has(msg.Error)
	}
	for _, errMsg := range r.ErrMessage {
		if errMsg != nil {
			errH.Has(errMsg)
		}
	}
	for _, decErr := range r.ErrMessageDecErr {
		errH.HasStr(decErr)
	}
	return errH.AllStr("\t")
}

func (r *EventResp) Err() error {
	if errStr := r.Error(); len(errStr) > 0 {
		return errors.New(errStr)
	}
	return nil
}

func (r *EventResp) err() bool {
	for _, msg := range r.Message {
		if msg != nil && msg.Error != nil {
			return true
		}
	}
	return len(r.ErrMessage) > 0 || len(r.ErrMessageDecErr) > 0
}
