package etcdv3

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
	etcdcliv3 "go.etcd.io/etcd/client/v3"
)

type Client struct {
	cli *etcdcliv3.Client
	ctx context.Context
	dto int
}

func New(addr ...string) (*Client, error) {
	ctx := util.BackgroundContext()
	dto := DefaultSecond
	cli, err := etcdcliv3.New(
		etcdcliv3.Config{
			Endpoints:            addr,
			DialTimeout:          util.TimeSecond(dto),
			DialKeepAliveTime:    util.TimeSecond(dto),
			DialKeepAliveTimeout: util.TimeSecond(dto),
			Context:              ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		cli: cli,
		ctx: ctx,
	}, nil
}

const DefaultSecond = 8

func (c *Client) Close() error {
	return c.cli.Close()
}

func (c *Client) Get(k string) (resp *etcdcliv3.GetResponse, err error) {
	c.timeoutOp(func(ctx context.Context) {
		resp, err = c.cli.Get(ctx, k)
	})
	return
}

func (c *Client) Put(k, v string) (resp *etcdcliv3.PutResponse, err error) {
	c.timeoutOp(func(ctx context.Context) {
		resp, err = c.cli.Put(ctx, k, v)
	})
	return
}

func (c *Client) timeoutOp(op func(ctx context.Context)) {
	ctx, cancel := context.WithTimeout(c.ctx, util.TimeSecond(c.dto))
	op(ctx)
	cancel()
}
