package harbor_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/harbor"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/repository"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/user"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"net/http"
	"strings"
)

// move it to yaml config file
const (
	harborRegistry = "http://192.168.31.242:8662"
	harborUsername = "hln"
	harborPassword = "Hln@001202"
	harborPublic   = "public"
)

const (
	ProjectCallBack       CallBackType = "Project"
	RepoCallBack          CallBackType = "Repository"
	ProjectNoStorageLimit int64        = -1
)

var (
	todoContext       = context.TODO()
	defaultHttpClient = http.DefaultClient
)

type (
	Client struct {
		// global variant
		v2Cli      *v2client.HarborAPI
		ctx        context.Context
		httpCli    *http.Client
		pageConfig *util.Page

		// callback variant
		CallBack *CallBack
	}

	CallBackType      string
	CallBackOpt       string
	CallBackAttribute struct {
		latest any
		doFunc util.Func
	}
	CallBackAttr  map[CallBackOpt]CallBackAttribute
	CallBackAttrs map[CallBackType]CallBackAttr
	CallBack      struct {
		CallBackAttrs
	}

	ProjectReqConfig struct {
		MetaDataPublic string
		ProjectName    string
		RegistryId     int64
		StorageLimit   int64
	}

	ArtifactURI struct {
		Project    string
		Repository string
		Tag        string
	}
)

func New() (*Client, error) {
	var (
		csc = &harbor.ClientSetConfig{
			URL:      harborRegistry,
			Insecure: true,
			Username: harborUsername,
			Password: harborPassword,
		}
	)
	c := &Client{
		ctx:        todoContext,
		httpCli:    defaultHttpClient,
		pageConfig: util.DefaultPage,
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

// Example:
// library/s3fstest:latest
func (a ArtifactURI) String() string {
	return strings.Join([]string{
		a.Project, "/",
		a.Repository, ":",
		a.Tag,
	}, "")
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

func (c *Client) WithCallBack() *Client {
	return c
}

func (c *Client) ListProjects() (*project.ListProjectsOK, error) {
	return c.v2Cli.Project.ListProjects(
		c.ctx,
		project.NewListProjectsParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithPublic(util.NewBool(true)),
	)
}

// CreateProject
// preserve a project named: public
// create new project for each user by a specified name
// restrict the quota of each user
func (c *Client) CreateProject(pReqCfg ProjectReqConfig) (*project.CreateProjectCreated, error) {
	return c.v2Cli.Project.CreateProject(
		c.ctx,
		project.NewCreateProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithProject(NewProjectReq(pReqCfg)),
	)
}

func (c *Client) DeleteProject(pName string) (*project.DeleteProjectOK, error) {
	return c.v2Cli.Project.DeleteProject(
		c.ctx,
		project.NewDeleteProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithProjectNameOrID(pName),
	)
}

// listAllRepositories
// WARNING: high privilege and vague api: do not expose to Web
func (c *Client) listAllRepositories() (*repository.ListAllRepositoriesOK, error) {
	return c.v2Cli.Repository.ListAllRepositories(
		c.ctx,
		repository.NewListAllRepositoriesParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()),
	)
}

func (c *Client) ListRepositories(pName string) (*repository.ListRepositoriesOK, error) {
	return c.v2Cli.Repository.ListRepositories(
		c.ctx,
		repository.NewListRepositoriesParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithProjectName(pName),
	)
}

func (c *Client) ListArtifacts(arti ArtifactURI) (*artifact.ListArtifactsOK, error) {
	return c.v2Cli.Artifact.ListArtifacts(
		c.ctx,
		artifact.NewListArtifactsParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithProjectName(arti.Project).
			WithRepositoryName(arti.Repository).
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithWithTag(util.NewBool(true)),
	)
}

func (c *Client) CopyArtifact(toArti, fromArti ArtifactURI) (*artifact.CopyArtifactCreated, error) {
	return c.v2Cli.Artifact.CopyArtifact(
		c.ctx,
		artifact.NewCopyArtifactParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithProjectName(toArti.Project).
			WithRepositoryName(toArti.Repository).
			WithFrom(fromArti.String()),
	)
}

func (c *Client) GetArtifact(arti ArtifactURI) (*artifact.GetArtifactOK, error) {
	return c.v2Cli.Artifact.GetArtifact(
		c.ctx,
		artifact.NewGetArtifactParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithProjectName(arti.Project).
			WithRepositoryName(arti.Repository).
			WithReference(arti.Tag).
			WithWithTag(util.NewBool(true)),
	)
}

func (c *Client) DeleteArtifact(arti ArtifactURI) (*artifact.DeleteArtifactOK, error) {
	return c.v2Cli.Artifact.DeleteArtifact(
		c.ctx,
		artifact.NewDeleteArtifactParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithProjectName(arti.Project).
			WithRepositoryName(arti.Repository).
			WithReference(arti.Tag),
	)
}

func (c *Client) CreateArtifactTag(toArti, fromArti ArtifactURI) (*artifact.CreateTagCreated, error) {
	return c.v2Cli.Artifact.CreateTag(
		c.ctx,
		artifact.NewCreateTagParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithProjectName(fromArti.Project).
			WithRepositoryName(fromArti.Repository).
			WithReference(fromArti.Tag).
			WithTag(&models.Tag{
				Immutable: false,
				Name:      toArti.Tag,
			}),
	)
}

// ExportArtifact
// Download the image.tar.gz file to local
func (c *Client) ExportArtifact() (string, error) {

	return "", nil
}

// ImportOfflineArtifact
// Upload the image.tar.gz file to project
func (c *Client) ImportOfflineArtifact() error {

	return nil
}

// GenerateArtifact
// generate image from Dockerfile and push it to project
func (c *Client) GenerateArtifact() error {

	return nil
}

func (c *Client) CreateAdmin(userReq *models.UserCreationReq) (*user.CreateUserCreated, error) {
	return c.v2Cli.User.CreateUser(
		c.ctx,
		user.NewCreateUserParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithUserReq(userReq),
	)
}

func (c *Client) DeleteAdmin(uid int64) (*user.DeleteUserOK, error) {
	return c.v2Cli.User.DeleteUser(
		c.ctx,
		user.NewDeleteUserParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithUserID(uid),
	)
}
