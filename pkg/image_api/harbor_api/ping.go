package harbor_api

import "github.com/goharbor/go-client/pkg/sdk/v2.0/client/ping"

func (c *Client) Ping() (*ping.GetPingOK, error) {
	return c.v2Cli.Ping.GetPing(
		c.ctx,
		ping.NewGetPingParams().
			WithContext(c.ctx).
			WithTimeout(c.httpTimeout).
			WithHTTPClient(c.httpCli))
}
