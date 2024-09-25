package docker_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/pkg/jsonmessage"
	"io"
	"strings"
)

var (
	errImageNotFound = errors.New("image not found")
)

type ImageRef struct {
	Registry   string
	Project    string
	Repository string
	Tag        string

	// cache
	absRefStr string
}

func ParseImageRef(absRefStr string) (imageRef ImageRef) {
	refParts := strings.Split(absRefStr, "/")
	if len(refParts) != 3 {
		stdlog.ErrorF("absRefStr: %s do not contain 2 slash, parse error", absRefStr)
		return
	}
	imageRef.absRefStr = absRefStr
	imageRef.Registry = refParts[0]
	imageRef.Project = refParts[1]

	artifactParts := strings.Split(refParts[2], ":")
	if len(artifactParts) != 2 {
		stdlog.ErrorF("absRefStr: %s artifact do not contain colon, parse error", absRefStr)
		return
	}
	imageRef.Repository = artifactParts[0]
	imageRef.Tag = artifactParts[1]
	return
}

func (r *ImageRef) String() string {
	if len(r.absRefStr) == 0 {
		r.absRefStr = util.StringJoin("/", r.Registry, r.Project, r.Repository+":"+r.Tag)
	}
	return r.absRefStr
}

// ExportImage
// +param absRefStr
// absolutely reference string for registry/project/repository/artifact/name:tag
// +desc
//
//	(1).if image do not exist in docker host local, pull image to local
//
//	(2).get image id by absRefStr
//
//	(3).save image to docker host local get io.ReadCloser fd
func (c *Client) ExportImage(absRefStr string) (io.ReadCloser, error) {
	_, err := c.pullImage(absRefStr)
	if err != nil {
		return nil, err
	}

	imageList, err := c.listImageByRef(absRefStr)
	if err != nil {
		return nil, err
	}
	if len(imageList) == 0 {
		return nil, errImageNotFound
	}

	return c.cli.ImageSave(c.ctx, []string{imageList[0].ID})
}

// ImportImage
// +param input
//
// (1). make sure the input file(.tar.gz) metadata attribute image name format is: registry/project/image_name:image_tag
//
//	(1.1). or will create new image tag to local
//
// (2). push image to registry
func (c *Client) ImportImage(absRefStr string, input io.Reader) (io.ReadCloser, error) {
	loadResp, err := c.cli.ImageLoad(c.ctx, input, false)
	if err != nil {
		return nil, err
	}
	defer util.HandleCloseError("docker image load", loadResp.Body)

	if loadResp.Body != nil && loadResp.JSON {
		// return json message
		stdlog.Debug("docker image loadResp format: json")
		ignoreJSONMessageErr := jsonmessage.DisplayJSONMessagesToStream(loadResp.Body, stdlog.Stream(), nil)
		if ignoreJSONMessageErr != nil {
			stdlog.Warn(ignoreJSONMessageErr)
			stdlog.Info(util.IOGetStr(loadResp.Body))
		}
	} else {
		stdlog.Debug("docker image loadResp format: plain text")
		_, err = io.Copy(stdlog.Stream(), loadResp.Body)
		if err != nil {
			return nil, err
		}
	}

	// get Loaded image: name:tag
	var loadedImageRefStr string
	return c.CreateImageTag(absRefStr, loadedImageRefStr)

}

func (c *Client) CreateImageTag(toAbsRefStr, fromAsbRefStr string) (io.ReadCloser, error) {
	err := c.cli.ImageTag(c.ctx, fromAsbRefStr, toAbsRefStr)
	if err != nil {
		return nil, err
	}
	return c.pushImage(toAbsRefStr)
}

type HostImageGCFunc func()

// HostImageGC
// cli: docker rmi IMAGE_ID
// maybe quota by:
//
// 1. image CREATED: since, before
// 2. image SIZE: bytes(B)
// 3. cache algorithm policy: lru, lfu
// 4. host disk: bytes(B)
func (c *Client) HostImageGC(hostImageGCFunc HostImageGCFunc) {

}

func (c *Client) pullImage(absRefStr string) (io.ReadCloser, error) {
	return c.cli.ImagePull(c.ctx, absRefStr, image.PullOptions{All: false})
}

func (c *Client) pushImage(absRefStr string) (io.ReadCloser, error) {
	return c.cli.ImagePush(c.ctx, absRefStr, image.PushOptions{All: false})
}

func (c *Client) listImageByRef(absRefStr string) ([]image.Summary, error) {
	return c.cli.ImageList(c.ctx, image.ListOptions{
		All:            false,
		Filters:        getRefFilter(absRefStr),
		SharedSize:     false,
		ContainerCount: false,
		Manifests:      false,
	})
}

func (c *Client) searchImageByRef(absRefStr string) ([]registry.SearchResult, error) {
	return c.cli.ImageSearch(c.ctx, "", registry.SearchOptions{
		RegistryAuth:  "",
		PrivilegeFunc: nil,
		Filters:       getRefFilter(absRefStr),
	})
}

func getRefFilter(absRefStr string) filters.Args {
	if !validAbsRefStr(absRefStr) {
		return filters.NewArgs(filters.KeyValuePair{
			Key:   FilterReference,
			Value: ReferenceNone,
		})
	}
	return filters.NewArgs(filters.KeyValuePair{
		Key:   FilterReference,
		Value: absRefStr,
	})
}

func validAbsRefStr(absRefStr string) bool {
	return strings.Count(absRefStr, "/") == 2
}

// Deprecated
// +example +strict url format
// 192.168.31.242:8662/library/ubuntu-s:22.04
func getRelativeRefStr(absRefStr string) string {
	parts := strings.Split(absRefStr, "/")
	if len(parts) != 3 {
		stdlog.ErrorF("parse reference string %s do not contain 2 slash", absRefStr)
		return ""
	}

	return util.StringConcat(parts[1], "/", parts[2])
}

type (
	engineAPIv1dot43ImagesLoadResp struct {
		//{
		//	"Id": "sha256:abcdef123456...",
		//	"RepoTags": [
		//	"myimage:latest"
		//],
		//	"Message": ""
		//}
		Id          string   `json:"Id,omitempty"`
		RepoTags    []string `json:"RepoTags,omitempty"`
		Message     string   `json:"Message,omitempty"`
		Error       string   `json:"error,omitempty"`
		ErrorDetail any      `json:"errorDetail,omitempty"`
	}
)
