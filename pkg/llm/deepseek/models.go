package deepseek

import (
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func (c *Client) Models() (resp openai.ModelsList, err error) {
	rresp, err := c.cli.R().
		Get("/models")
	if err != nil {
		return
	} else if rresp.StatusCode() != http.StatusOK {
		err = ErrRespNotOK
		return
	}
	err = Dec(rresp.Body(), &resp)
	return
}
