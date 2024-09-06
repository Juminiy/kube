package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"strings"
)

type ArtifactURI struct {
	Project    string
	Repository string
	Tag        string
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

func (c *Client) ListArtifacts(artifactURI ArtifactURI) (*artifact.ListArtifactsOK, error) {
	return c.v2Cli.Artifact.ListArtifacts(
		c.ctx,
		artifact.NewListArtifactsParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
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
			WithProjectName(artifactURI.Project).
			WithRepositoryName(artifactURI.Repository).
			WithReference(artifactURI.Tag),
	)
}

func (c *Client) CreateArtifactTag(toArtifactURI, fromArtifactURI ArtifactURI) (*artifact.CreateTagCreated, error) {
	return c.v2Cli.Artifact.CreateTag(
		c.ctx,
		artifact.NewCreateTagParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithProjectName(fromArtifactURI.Project).
			WithRepositoryName(fromArtifactURI.Repository).
			WithReference(fromArtifactURI.Tag).
			WithTag(&models.Tag{
				Immutable: false,
				Name:      toArtifactURI.Tag,
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
