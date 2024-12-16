package docker_client

import (
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/registry"
	"strings"
)

var errRefIsDigest = errors.New("cannot push a digest reference")

func (c *Client) ImagePush(refStr string) (pushResp EventResp, err error) {
	ref, err := reference.ParseNormalizedNamed(refStr)
	if err != nil {
		return
	}
	if _, isCanonical := ref.(reference.Canonical); isCanonical {
		return pushResp, errRefIsDigest
	}

	r := c.post("/images/{name}/push")
	name := reference.FamiliarName(ref)
	ref = reference.TagNameOnly(ref)
	if tagged, ok := ref.(reference.Tagged); ok {
		r.r.SetQueryParam("tag", tagged.Tag())
	}
	r.r.SetPathParam("name", name).
		SetHeader(registry.AuthHeader, c.reg.Auth)

	return pushResp.WrapParse(r.do())
}

type ImagePushResp struct {
	RemoteRepository string
	PushID           string
	Aux              struct {
		Tag    string
		Digest string
		Size   int
	}
}

func (r *EventResp) GetImagePushResp() (resp ImagePushResp) {
	for _, msg := range r.Message {
		if msg == nil {
			continue
		}
		if msg.Aux != nil {
			err := DecE(*msg.Aux, &resp.Aux)
			if err != nil {
				stdlog.ErrorF("unmarshal image aux error: %s", err.Error())
			}
		}
		if statusStr := msg.Status; strings.HasPrefix(statusStr, "The push refers to repository ") {
			statusStr = strings.TrimPrefix(statusStr, "The push refers to repository ")
			resp.RemoteRepository = util.StringDelete(statusStr, "[", "]")
		}
		if len(msg.ID) > 0 {
			resp.PushID = msg.ID
		}
	}
	return
}
