package docker_api

import "github.com/docker/docker/api/types"

func (c *Client) SystemVersion() (types.Version, error) {
	return c.cli.ServerVersion(c.ctx)
}
