package docker_client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
	"sync"
)

type req struct {
	method string
	url    string
	r      *resty.Request
	o      sync.Once
	c      *Client
}

func (r *req) do() (resp *resty.Response, err error) {
	r.o.Do(func() {
		resp, err = r.r.Execute(r.method, fmt.Sprintf("/%s/%s", r.c.version, strings.TrimPrefix(r.url, "/")))
	})
	return
}

func (c *Client) req(method string, url string) *req {
	return &req{
		method: method,
		url:    url,
		r:      c.rCli.R(),
		o:      sync.Once{},
		c:      c,
	}
}

func (c *Client) get(url string) *req {
	return c.req(http.MethodGet, url)
}

func (c *Client) post(url string) *req {
	return c.req(http.MethodPost, url)
}

func (c *Client) put(url string) *req {
	return c.req(http.MethodPut, url)
}

func (c *Client) patch(url string) *req {
	return c.req(http.MethodPatch, url)
}

func (c *Client) delete(url string) *req {
	return c.req(http.MethodDelete, url)
}
