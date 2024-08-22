package harbor_api

import (
	"context"
	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"kube/util"
	"net/http"
)

// move it to yaml
const (
	harborRegistry = "192.168.31.242:8662"
	harborUsername = "admin"
	harborPassword = "Harbor12345"
)

type Client struct {
	Name string
	csc  *harbor.ClientSetConfig
	cli  *harbor.ClientSet
	ctx  context.Context
}

func NewRepoClient(repoName string) (*Client, error) {
	var err error
	c := &Client{
		Name: repoName,
		csc: &harbor.ClientSetConfig{
			URL:      harborRegistry,
			Insecure: true,
			Username: harborUsername,
			Password: harborPassword,
		},
		ctx: context.TODO(),
	}
	c.cli, err = harbor.NewClientSet(c.csc)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c Client) ListImages() (*project.ListProjectsOK, error) {
	return c.cli.V2().Project.ListProjects(
		c.ctx,
		project.NewListProjectsParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithName(util.NewString(c.Name)).
			WithPublic(util.NewBool(true)),
	)
}
