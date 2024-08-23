package harbor_api

import (
	"context"
	"github.com/goharbor/go-client/pkg/harbor"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/repository"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"kube/util"
	"net/http"
)

// move it to yaml
const (
	harborRegistry = "http://192.168.31.242:8662"
	harborUsername = "hln"
	harborPassword = "Hln@001202"
)

const (
	ProjectCallBack CallBackType = "Project"
	RepoCallBack    CallBackType = "Repository"
)

type (
	Client struct {
		// global variant
		PageConfig *util.Page
		cliV2      *v2client.HarborAPI
		ctx        context.Context

		// project variant
		ProjectReqConfig

		TriggerCallBack bool
		CallBack
	}

	CallBackType      string
	CallBackOpt       string
	CallBackAttribute struct {
		latest any
		doFunc util.Func
	}
	CallBack map[CallBackType]map[CallBackOpt]CallBackAttribute

	ProjectReqConfig struct {
		MetaDataPublic string
		ProjectName    string
		RegistryId     int64
		StorageLimit   int64
	}
)

func NewHarborCli() (*Client, error) {
	var (
		csc = &harbor.ClientSetConfig{
			URL:      harborRegistry,
			Insecure: true,
			Username: harborUsername,
			Password: harborPassword,
		}
	)
	c := &Client{
		ctx: context.TODO(),
	}
	hCli, err := harbor.NewClientSet(csc)
	if err != nil {
		return nil, err
	}
	c.cliV2 = hCli.V2()
	return c, nil
}

var (
	defaultHttpClient = http.DefaultClient
)

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

func (c *Client) WithProjectName(pName string) *Client {
	c.ProjectName = pName
	return c
}

func (c *Client) WithPageConfig(pCfg *util.Page) *Client {
	c.PageConfig = pCfg
	return c
}

func (c *Client) WithCallBack() *Client {
	c.TriggerCallBack = true
	return c
}

func (c *Client) ListProjects() (*project.ListProjectsOK, error) {
	return c.cliV2.Project.ListProjects(
		c.ctx,
		project.NewListProjectsParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			//WithName(util.NewString(c.ProjectName)).
			WithPublic(util.NewBool(true)),
	)
}

// CreateProject
// preserve a project named: public
// create new project for each user by a specified name
// restrict the quota of each user
func (c *Client) CreateProject() (*project.CreateProjectCreated, error) {
	return c.cliV2.Project.CreateProject(
		c.ctx,
		project.NewCreateProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithProject(NewProjectReq(c.ProjectReqConfig)),
	)
}

func (c *Client) DeleteProject() (*project.DeleteProjectOK, error) {
	return c.cliV2.Project.DeleteProject(
		c.ctx,
		project.NewDeleteProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithProjectNameOrID(c.ProjectName),
	)
}

// WARNING: high privilege and vague api: do not expose to Web
func (c *Client) listAllImages() (*repository.ListAllRepositoriesOK, error) {
	return c.cliV2.Repository.ListAllRepositories(
		c.ctx,
		repository.NewListAllRepositoriesParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithPage(c.PageConfig.Page()).
			WithPageSize(c.PageConfig.Size()),
	)
}

func (c *Client) ListImages() (*repository.ListRepositoriesOK, error) {
	return c.cliV2.Repository.ListRepositories(
		c.ctx,
		repository.NewListRepositoriesParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithPage(c.PageConfig.Page()).
			WithPageSize(c.PageConfig.Size()).
			WithProjectName(c.ProjectName),
	)
}

// Download the image.tar.gz file to localdir
func (c *Client) GetImageDLURL() (string, error) {

	return "", nil
}

// Upload a image.tar.gz file to repo
func (c *Client) UploadImage() error {

	return nil
}

// generate image from Dockerfile and push it to
func (c *Client) GenerateImage() error {

	return nil
}

func (c *Client) DeleteImage() error {

	return nil
}
