package docker_api

import "github.com/docker/docker/api/types"

func (c *Client) SystemPing() (types.Ping, error) {
	return c.cli.Ping(c.ctx)
}
