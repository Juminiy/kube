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

	resp, err := r.do()
	return *loadResp.Parse(resp), err
}

/*
case1: refStr

	{
	    "stream": "Loaded image: hello-world:latest\n"
	}

case2: imageID

	{
		"stream": "Loaded image ID: sha256:d2c94e258dcb3c5ac2798d32e1249e42ef01cba4841c2234249495f87264ac5a\n"
	}
*/
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
