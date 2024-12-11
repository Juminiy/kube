package docker_client

import (
	"bytes"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/go-resty/resty/v2"
)

var EncE = safe_json.Goccy().Marshal
var DecE = safe_json.Goccy().Unmarshal

type EventResp struct {
	Status     int
	Message    []*jsonmessage.JSONMessage
	MessageErr []error // unmarshal Message error
}

func (r *EventResp) Parse(resp *resty.Response) *EventResp {
	r.Status = resp.StatusCode()
	return r.parse(resp.Body())
}

func (r *EventResp) ParseBytes(bs []byte) *EventResp {
	return r.parse(bs)
}

func (r *EventResp) parse(bs []byte) *EventResp {
	if len(bs) == 0 {
		return r
	}
	sbs := bytes.Split(bs, []byte{'\r', '\n'})
	r.Message = make([]*jsonmessage.JSONMessage, len(sbs))
	r.MessageErr = make([]error, len(sbs))
	for i := range sbs {
		if len(sbs[i]) == 0 {
			continue
		}
		msg := jsonmessage.JSONMessage{}
		err := DecE(sbs[i], &msg)
		if err == nil {
			r.Message[i] = &msg
		} else {
			r.MessageErr[i] = err
		}
	}
	return r
}
