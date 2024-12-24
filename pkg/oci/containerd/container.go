package containerd

import (
	containerdcli "github.com/containerd/containerd/v2/client"
)

func (c *Client) ContainerList(filters ...string) ([]containerdcli.Container, error) {
	return c.cli.Containers(c.ctx, filters...)
}
