package docker_client

import (
	"io"
	"strings"
)

func (c *Client) ImageLoad(input io.Reader) (loadResp EventResp, err error) {
	r := c.post("/images/load")
	r.r.SetHeader("Content-Type", "application/x-tar").
		SetQueryParam("quiet", "1").
		SetBody(input)

	return loadResp.WrapParse(r.do())
}

func (r *EventResp) GetImageLoad() (refStrOrImageID string) {
	for _, msg := range r.Message {
		if msg == nil {
			continue
		}
		if strings.HasPrefix(msg.Stream, "Loaded image: ") {
			refStrOrImageID = strings.TrimPrefix(msg.Stream, "Loaded image: ")
			return strings.TrimSuffix(refStrOrImageID, "\n")
		} else if strings.HasPrefix(msg.Stream, "Loaded image ID: ") {
			refStrOrImageID = strings.TrimPrefix(msg.Stream, "Loaded image ID: ")
			return strings.TrimSuffix(refStrOrImageID, "\n")
		}
	}
	return
}
