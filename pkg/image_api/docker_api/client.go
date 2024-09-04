package docker_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	dockercli "github.com/docker/docker/client"
)

const (
	hostURL = "tcp://192.168.31.242:2376"
)

type Client struct {
	cli *dockercli.Client
	ctx context.Context
}

func New() *Client {
	dCli, err := dockercli.NewClientWithOpts(
		dockercli.WithHost(hostURL),
	)
	if err != nil {
		stdlog.ErrorF("connect to docker host: %s error: %s", hostURL, err.Error())
		return nil
	}
	return &Client{
		cli: dCli,
		ctx: context.TODO(),
	}
}
