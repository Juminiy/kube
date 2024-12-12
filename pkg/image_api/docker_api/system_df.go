package docker_api

import "github.com/docker/docker/api/types"

func (c *Client) SystemDF(objectTypes ...types.DiskUsageObject) (types.DiskUsage, error) {
	return c.cli.DiskUsage(c.ctx, types.DiskUsageOptions{
		Types: objectTypes,
	})
}
