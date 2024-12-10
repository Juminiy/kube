package docker_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_client"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/errdefs"
	"io"
)

type ExportImageResp struct {
	RequestRefStr      string
	NotFoundInRegistry bool // absRefStr not found in provided registry.
	NotFoundInLocal    bool // absRefStr not found in docker-server local.
	// deprecated in Client.ExportImageV2
	ImagePulledInfo string        // image pulled information from provided registry.
	ImageFileReader io.ReadCloser `json:"-"` // compressed (.tar, .tar.gz) file reader, It's up to the caller to handle the io.ReadCloser and close it properly.
}

// ExportImage
// +param absRefStr
// (1). absolutely reference string for registry/project/repository/artifact/name:tag
// +desc
//
//	(1).if image do not exist in docker host local, pull image to local
//
//	(2).get image id by absRefStr
//
//	(3).save image to docker host local get io.ReadCloser fd
func (c *Client) ExportImage(absRefStr string) (resp ExportImageResp, err error) {
	resp.RequestRefStr = absRefStr
	// always pull image from registry
	imagePulled, err := c.pullImage(absRefStr)
	if err != nil {
		switch err.(type) {
		case errdefs.ErrSystem:
			resp.NotFoundInRegistry = docker_internal.GetErrParser().ArtifactNotFound(err)
		default:
			return resp, err
		}
	}
	if imagePulled != nil {
		defer util.SilentCloseIO("readCloser", imagePulled)
		resp.ImagePulledInfo = util.IOGetStr(imagePulled)
	}

	// inspect image id, lookup if in server local
	imageInspect, err := c.InspectImage(absRefStr)
	if err != nil {
		switch err.(type) {
		case errdefs.ErrNotFound:
			resp.NotFoundInLocal = true
		}
		return resp, err
	}

	resp.ImageFileReader, err = c.cli.ImageSave(c.ctx, []string{imageInspect.ID})
	return resp, err
}

type ExportImageRespV2 struct {
	ExportImageResp
	docker_client.ImagePullResp
}

var ErrNotFoundImageAnyWhere = errors.New("not found image refStr in anywhere")

func (c *Client) ExportImageV2(absRefStr string) (resp ExportImageRespV2, err error) {
	resp.RequestRefStr = absRefStr

	pullResp, err := c.apiClient.ImagePull(absRefStr)
	if err != nil {
		return
	}
	resp.ImagePullResp = pullResp.GetImagePullResp()
	resp.NotFoundInRegistry = resp.ImagePullResp.NotFound

	inspect, err := c.InspectImage(absRefStr)
	if err != nil {
		switch err.(type) {
		case errdefs.ErrNotFound:
			resp.NotFoundInLocal = true
		default:
			return resp, err
		}
	}
	if resp.NotFoundInLocal && resp.NotFoundInRegistry {
		return resp, ErrNotFoundImageAnyWhere
	}

	resp.ImageFileReader, err = c.cli.ImageSave(c.ctx, []string{inspect.ID})
	return
}
