package containerd

import (
	"github.com/Juminiy/kube/pkg/util"
	containerdcli "github.com/containerd/containerd/v2/client"
)

func (c *Client) ImageList(filters ...string) ([]containerdcli.Image, error) {
	images, err := c.cli.ListImages(c.ctx, filters...)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (c *Client) ImagePull(absRef string) (containerdcli.Image, error) {
	image, err := c.cli.Pull(c.ctx, absRef,
		func(cli *containerdcli.Client, rctx *containerdcli.RemoteContext) error {
			rctx.Unpack = true
			rctx.MaxConcurrentDownloads = util.MagicNumber
			rctx.MaxConcurrentUploadedLayers = util.MagicNumber
			return nil
		})
	if err != nil {
		return nil, err
	}
	return image, nil
}
