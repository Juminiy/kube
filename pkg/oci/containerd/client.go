package containerd

import (
	"context"

	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/util"
	containerdcli "github.com/containerd/containerd/v2/client"
	"github.com/containerd/platforms"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

type Client struct {
	cli *containerdcli.Client

	ctx  context.Context
	page *util.Page
}

func New() (*Client, error) {
	cli, err := containerdcli.New(DefaultAddr,
		containerdcli.WithDefaultPlatform(DefaultPlatform),
		containerdcli.WithDefaultNamespace(DefaultNS),
		containerdcli.WithTimeout(util.TimeSecond(DefaultTOSec)),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		cli:  cli,
		ctx:  util.TODOContext(),
		page: util.DefaultPage(),
	}, nil
}

const DefaultAddr = "/run/containerd/containerd.sock"
const DefaultNS = "LYY"
const DefaultTOSec = 8

var DefaultPlatform = platforms.Only(specs.Platform{
	Architecture: internal_api.Amd64,
	OS:           internal_api.Linux,
	OSVersion:    "",
	OSFeatures:   nil,
	Variant:      "",
})

func (c *Client) Server() (ServerUUID string, err error) {
	serverInfo, err := c.cli.Server(c.ctx)
	if err != nil {
		return "", err
	}
	return serverInfo.UUID, err
}
