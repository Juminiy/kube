package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
)

func (c *Client) ErrorDetail(err error) string {
	switch ev := err.(type) {
	case *project.CreateProjectBadRequest:
		bs, err := ev.Payload.MarshalBinary()
		if err != nil {
			return err.Error()
		}
		return util.Bytes2StringNoCopy(bs)
	}

	return ""
}
