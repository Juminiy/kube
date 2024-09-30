// Package harbor_api
package harbor_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/harbor"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"net/http"
)

type Client struct {
	// global variant
	v2Cli      *v2client.HarborAPI
	ctx        context.Context
	httpCli    *http.Client
	pageConfig *util.Page

	// callback variant
	CallBack *CallBack
}

func New(
	harborRegistry string,
	harborInsecure bool,
	harborUsername,
	harborPassword string) (*Client, error) {
	csc := &harbor.ClientSetConfig{
		URL:      harborRegistry,
		Insecure: harborInsecure,
		Username: harborUsername,
		Password: harborPassword,
	}
	c := &Client{
		ctx:        util.TODOContext(),
		httpCli:    util.DefaultHTTPClient(),
		pageConfig: util.DefaultPage(),
	}
	hCli, err := harbor.NewClientSet(csc)
	if err != nil {
		return nil, err
	}
	c.v2Cli = hCli.V2()
	return c, nil
}

func NewProjectReq(reqCfg ProjectReqConfig) *models.ProjectReq {
	return &models.ProjectReq{
		Metadata:     &models.ProjectMetadata{Public: reqCfg.MetaDataPublic},
		ProjectName:  reqCfg.ProjectName,
		RegistryID:   util.NewInt64(reqCfg.RegistryId),
		StorageLimit: util.NewInt64(reqCfg.StorageLimit),
	}
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

func (c *Client) WithHttpClient(httpCli *http.Client) *Client {
	c.httpCli = httpCli
	return c
}

func (c *Client) WithPageConfig(pCfg *util.Page) *Client {
	c.pageConfig = pCfg
	return c
}

func (c *Client) WithCallBack(callback *CallBack) *Client {
	c.CallBack = callback
	return c
}

func (c *Client) GC(gcFn ...util.Func) {

}
