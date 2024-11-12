package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

const (
	ProjectNoStorageLimit int64 = -1
)

type ProjectReqConfig struct {
	MetaDataPublic string
	ProjectName    string
	RegistryId     int64
	StorageLimit   int64
}

func NewProjectReq(reqCfg ProjectReqConfig) *models.ProjectReq {
	return &models.ProjectReq{
		Metadata:     &models.ProjectMetadata{Public: reqCfg.MetaDataPublic},
		ProjectName:  reqCfg.ProjectName,
		RegistryID:   util.NewNoZeroInt(reqCfg.RegistryId),
		StorageLimit: util.NewNoZeroInt(reqCfg.StorageLimit),
	}
}

// CreateProject
// preserve a project named: public
// create new project for each user by a specified name
// restrict the quota of each user
func (c *Client) CreateProject(projectReqConfig ProjectReqConfig) (*project.CreateProjectCreated, error) {
	return c.v2Cli.Project.CreateProject(
		c.ctx,
		project.NewCreateProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithDefaults().
			WithProject(NewProjectReq(projectReqConfig)),
	)
}

func (c *Client) DeleteProject(projectNameOrID string) (*project.DeleteProjectOK, error) {
	return c.v2Cli.Project.DeleteProject(
		c.ctx,
		project.NewDeleteProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithDefaults().
			WithProjectNameOrID(projectNameOrID),
	)
}

func (c *Client) GetProject(projectNameOrID string) (*project.GetProjectOK, error) {
	return c.v2Cli.Project.GetProject(
		c.ctx,
		project.NewGetProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithDefaults().
			WithProjectNameOrID(projectNameOrID),
	)
}

func (c *Client) UpdateProjectStorageLimit(reqCfg *ProjectReqConfig) (*project.UpdateProjectOK, error) {
	return c.v2Cli.Project.UpdateProject(
		c.ctx,
		project.NewUpdateProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectNameOrID(reqCfg.ProjectName).
			WithProject(NewProjectReq(*reqCfg)),
	)
}

func (c *Client) ListProjects(public bool) (*project.ListProjectsOK, error) {
	return c.v2Cli.Project.ListProjects(
		c.ctx,
		project.NewListProjectsParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithDefaults().
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithPublic(util.NewBool(public)),
	)
}
