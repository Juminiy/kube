package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

type ArtifactURI struct {
	Project    string
	Repository string
	Tag        string
	Digest     string
}

// Example:
// library/s3fstest:latest
func (a ArtifactURI) String() string {
	if len(a.Tag) != 0 {
		return util.StringConcat(
			a.Project, "/",
			a.Repository, ":",
			a.Tag)
	} else if len(a.Digest) != 0 {
		return util.StringConcat(
			a.Project, "/",
			a.Repository, "@",
			a.Digest,
		)
	}
	return ""
}

func (c *Client) ListArtifacts(artifactURI ArtifactURI) (*artifact.ListArtifactsOK, error) {
	return c.v2Cli.Artifact.ListArtifacts(
		c.ctx,
		artifact.NewListArtifactsParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectName(artifactURI.Project).
			WithRepositoryName(artifactURI.Repository).
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithWithTag(util.NewBool(true)),
	)
}

func (c *Client) CopyArtifact(toArtifactURI, fromArtifactURI ArtifactURI) (*artifact.CopyArtifactCreated, error) {
	return c.v2Cli.Artifact.CopyArtifact(
		c.ctx,
		artifact.NewCopyArtifactParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectName(toArtifactURI.Project).
			WithRepositoryName(toArtifactURI.Repository).
			WithFrom(fromArtifactURI.String()),
	)
}

func (c *Client) GetArtifact(artifactURI ArtifactURI) (*artifact.GetArtifactOK, error) {
	return c.v2Cli.Artifact.GetArtifact(
		c.ctx,
		artifact.NewGetArtifactParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectName(artifactURI.Project).
			WithRepositoryName(artifactURI.Repository).
			WithReference(artifactURI.Tag).
			WithWithTag(util.NewBool(true)),
	)
}

func (c *Client) DeleteArtifact(artifactURI ArtifactURI) (*artifact.DeleteArtifactOK, error) {
	return c.v2Cli.Artifact.DeleteArtifact(
		c.ctx,
		artifact.NewDeleteArtifactParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectName(artifactURI.Project).
			WithRepositoryName(artifactURI.Repository).
			WithReference(artifactURI.Tag),
	)
}

func (c *Client) ListArtifactTags(artifactURI ArtifactURI, queryStr string) (*artifact.ListTagsOK, error) {
	return c.v2Cli.Artifact.ListTags(
		c.ctx,
		artifact.NewListTagsParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectName(artifactURI.Project).
			WithRepositoryName(artifactURI.Repository).
			WithReference(artifactURI.Tag).
			WithPage(c.pageConfig.Page()).
			WithPageSize(c.pageConfig.Size()).
			WithQ(util.NewString(queryStr)),
	)
}

func (c *Client) CreateArtifactTag(toArtifactURI, fromArtifactURI ArtifactURI) (*artifact.CreateTagCreated, error) {
	return c.v2Cli.Artifact.CreateTag(
		c.ctx,
		artifact.NewCreateTagParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectName(fromArtifactURI.Project).
			WithRepositoryName(fromArtifactURI.Repository).
			WithReference(fromArtifactURI.Tag).
			WithTag(&models.Tag{
				Immutable: false,
				Name:      toArtifactURI.Tag,
			}),
	)
}

func (c *Client) DeleteArtifactTag(artifactURI ArtifactURI) (*artifact.DeleteTagOK, error) {
	return c.v2Cli.Artifact.DeleteTag(
		c.ctx,
		artifact.NewDeleteTagParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithProjectName(artifactURI.Project).
			WithRepositoryName(artifactURI.Repository).
			WithReference(artifactURI.Tag),
	)
}

type ArtifactCopyTagGet struct {
	*artifact.CopyArtifactCreated
	*artifact.CreateTagCreated
	*artifact.GetArtifactOK
}

func (c *Client) ArtifactCopyTagGet(toURI, fromURI ArtifactURI) (resp ArtifactCopyTagGet, err error) {
	resp.CopyArtifactCreated, err = c.CopyArtifact(toURI, fromURI)
	if err != nil {
		return
	}

	copiedURI := ArtifactURI{
		Project:    toURI.Project,
		Repository: toURI.Repository,
		Tag:        fromURI.Tag,
	}
	copiedArtifact, err := c.GetArtifact(copiedURI)
	if err != nil {
		return
	}
	for _, tag := range copiedArtifact.Payload.Tags {
		if tag.Name == toURI.Tag {
			resp.GetArtifactOK = copiedArtifact
			return
		}
	}
	resp.CreateTagCreated, err = c.CreateArtifactTag(toURI, copiedURI)
	if err != nil {
		return
	}
	resp.GetArtifactOK, err = c.GetArtifact(toURI)
	return
}

func (c *Client) GC(gcFn ...util.Func) {

}
