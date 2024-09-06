package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
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
func (c *Client) CreateProject(projectReqConfig ProjectReqConfig) (*project.CreateProjectCreated, error) {
	return c.v2Cli.Project.CreateProject(
		c.ctx,
		project.NewCreateProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithProject(NewProjectReq(projectReqConfig)),
	)
}

func (c *Client) DeleteProject(projectName string) (*project.DeleteProjectOK, error) {
	return c.v2Cli.Project.DeleteProject(
		c.ctx,
		project.NewDeleteProjectParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithProjectNameOrID(projectName),
	)
}
