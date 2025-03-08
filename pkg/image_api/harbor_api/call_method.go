package harbor_api

import v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"

func (c *Client) Do(fc func(v2Cli *v2client.HarborAPI) error) error {
	return fc(c.v2Cli)
}
