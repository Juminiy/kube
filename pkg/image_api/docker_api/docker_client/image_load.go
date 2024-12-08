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
	{
	    "stream": "Loaded image: hello-world:latest\n"
	}
*/
func (r *EventResp) GetImageLoadRefStr() (refStr string) {
	for _, msg := range r.Message {
		if msg == nil {
			continue
		}
		if strings.HasPrefix(msg.Stream, "Loaded image: ") {
			refStr = strings.TrimPrefix(msg.Stream, "Loaded image: ")
			return strings.TrimSuffix(refStr, "\n")
		}
	}
	return
}
