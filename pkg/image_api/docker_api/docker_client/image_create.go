package docker_client

import (
	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/registry"
	"net/http"
	"strings"
)

func (c *Client) ImageCreate(parentRefStr string) (createResp EventResp, err error) {
	ref, err := reference.ParseNormalizedNamed(parentRefStr)
	if err != nil {
		return
	}

	return c.imageCreate(map[string]string{
		"fromImage": reference.FamiliarName(ref),
		"tag":       getAPITagFromNamedRef(ref),
	})
}

func (c *Client) imageCreate(queryParam map[string]string) (createResp EventResp, err error) {
	r := c.post("/images/create")
	r.r.SetQueryParams(queryParam).
		SetHeader(registry.AuthHeader, c.reg.Auth)

	return createResp.WrapParse(r.do())
}

func getAPITagFromNamedRef(ref reference.Named) string {
	if digested, ok := ref.(reference.Digested); ok {
		return digested.Digest().String()
	}
	ref = reference.TagNameOnly(ref)
	if tagged, ok := ref.(reference.Tagged); ok {
		return tagged.Tag()
	}
	return ""
}

type ImageCreateResp struct {
	PullFromRefStr  string
	DownLocalRefStr string
	Digest          string
	NotFound        bool
}

func (r *EventResp) GetImageCreateResp() (resp ImageCreateResp) {
	if r.Status == http.StatusNotFound {
		resp.NotFound = true
		return
	}
	for _, msg := range r.Message {
		if msg == nil {
			return
		}
		statusStr := msg.Status
		if strings.HasPrefix(statusStr, "Pulling from ") {
			resp.PullFromRefStr = strings.TrimPrefix(statusStr, "Pulling from ") + ":" + msg.ID
		}
		if strings.HasPrefix(statusStr, "Digest: ") {
			resp.Digest = strings.TrimPrefix(statusStr, "Digest: ")
		}
		if strings.HasPrefix(statusStr, "Status: Downloaded newer image for ") {
			resp.DownLocalRefStr = strings.TrimPrefix(statusStr, "Status: Downloaded newer image for ")
		} else if strings.HasPrefix(statusStr, "Status: Image is up to date for ") {
			resp.DownLocalRefStr = strings.TrimPrefix(statusStr, "Status: Image is up to date for ")
		}
	}
	return
}

/*
	httpStatusCode
	- 200 no error
	- 404 repository does not exist or no read access
	- 500 server error
*/
