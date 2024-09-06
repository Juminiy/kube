package harbor_api

import "github.com/goharbor/go-client/pkg/sdk/v2.0/client/repository"

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

func (c *Client) ListRepositories(projectName string) (*repository.ListRepositoriesOK, error) {
	return c.v2Cli.Repository.ListRepositories(
		c.ctx,
		repository.NewListRepositoriesParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithDefaults().
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithProjectName(projectName),
	)
}
