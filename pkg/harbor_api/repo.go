package harbor_api

import (
	"context"
	"github.com/goharbor/go-client/pkg/harbor"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/repository"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/user"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"kube/pkg/util"
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
		cliV2      *v2client.HarborAPI
		ctx        context.Context
		pageConfig *util.Page

		// callback variant
		TriggerCallBack bool
		CallBacks
	}

	CallBackType      string
	CallBackOpt       string
	CallBackAttribute struct {
		latest any
		doFunc util.Func
	}
	CallBack  map[CallBackOpt]CallBackAttribute
	CallBacks map[CallBackType]CallBack

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
		ctx:        context.TODO(),
		pageConfig: util.DefaultPage,
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

func (c *Client) WithPageConfig(pCfg *util.Page) *Client {
	c.pageConfig = pCfg
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
func (c *Client) CreateProject(pReqCfg ProjectReqConfig) (*project.CreateProjectCreated, error) {
	return c.cliV2.Project.CreateProject(
		c.ctx,
		project.NewCreateProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithProject(NewProjectReq(pReqCfg)),
	)
}

func (c *Client) DeleteProject(pName string) (*project.DeleteProjectOK, error) {
	return c.cliV2.Project.DeleteProject(
		c.ctx,
		project.NewDeleteProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithProjectNameOrID(pName),
	)
}

// WARNING: high privilege and vague api: do not expose to Web
func (c *Client) listAllRepositories() (*repository.ListAllRepositoriesOK, error) {
	return c.cliV2.Repository.ListAllRepositories(
		c.ctx,
		repository.NewListAllRepositoriesParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()),
	)
}

func (c *Client) ListRepositories(pName string) (*repository.ListRepositoriesOK, error) {
	return c.cliV2.Repository.ListRepositories(
		c.ctx,
		repository.NewListRepositoriesParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithDefaults().
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithProjectName(pName),
	)
}

func (c *Client) ListArtifacts(pName, rName string) (*artifact.ListArtifactsOK, error) {
	lsParam := artifact.NewListArtifactsParams().
		WithContext(c.ctx).
		WithHTTPClient(http.DefaultClient).
		WithProjectName(pName)
	lsParam.Page, lsParam.PageSize = c.pageConfig.Pair()
	lsParam.WithRepositoryName(rName)

	return c.cliV2.Artifact.ListArtifacts(
		c.ctx,
		lsParam,
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

func (c *Client) CreateAdmin(userReq *models.UserCreationReq) (*user.CreateUserCreated, error) {
	return c.cliV2.User.CreateUser(
		c.ctx,
		user.NewCreateUserParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithUserReq(userReq),
	)
}

func (c *Client) DeleteAdmin(uid int64) (*user.DeleteUserOK, error) {
	return c.cliV2.User.DeleteUser(
		c.ctx,
		user.NewDeleteUserParams().
			WithContext(c.ctx).
			WithHTTPClient(http.DefaultClient).
			WithUserID(uid),
	)
}
