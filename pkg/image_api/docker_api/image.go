package docker_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"io"
	"strings"
)

var (
	errImageNotFound = errors.New("image not found")
	errImageInternal = errors.New("image internal error")
)

// InspectImage
// +param image: sha256ID or refStr
func (c *Client) InspectImage(image string) (types.ImageInspect, error) {
	// discard raw []byte, no usage anymore
	imageInspect, _, err := c.cli.ImageInspectWithRaw(c.ctx, image)
	return imageInspect, err
}

func (c *Client) pullImage(absRefStr string) (io.ReadCloser, error) {
	return c.cli.ImagePull(c.ctx, absRefStr, image.PullOptions{
		All:          false,
		RegistryAuth: c.reg.Auth,
	})
}

func (c *Client) pushImage(absRefStr string) (io.ReadCloser, error) {
	return c.cli.ImagePush(c.ctx, absRefStr, image.PushOptions{
		All:          false,
		RegistryAuth: c.reg.Auth,
	})
}

// in docker host: list images by filter (like search)
// search images in local docker host
func (c *Client) listImageByRef(absRefStr string) ([]image.Summary, error) {
	return c.cli.ImageList(c.ctx, image.ListOptions{
		All:            false,
		Filters:        getRefFilter(absRefStr),
		SharedSize:     false,
		ContainerCount: false,
		Manifests:      false,
	})
}

// only in Dockerhub: search images by filter
func (c *Client) searchImageByRef(absRefStr string) ([]registry.SearchResult, error) {
	return c.cli.ImageSearch(c.ctx, "", registry.SearchOptions{
		RegistryAuth:  c.reg.Auth,
		PrivilegeFunc: nil,
		Filters:       getRefFilter(absRefStr),
	})
}

func getRefFilter(absRefStr string) (args filters.Args) {
	arti := ParseToArtifact(absRefStr)
	if refStr := arti.RefStr(); len(refStr) > 0 {
		return filters.NewArgs(filters.KeyValuePair{
			Key:   docker_internal.FilterReference,
			Value: refStr,
		})
	}
	return
}

// Deprecated
// +example +strict url format
// harbor.local:8080/library/ubuntu-s:22.04
func getRelativeRefStr(absRefStr string) string {
	parts := strings.Split(absRefStr, "/")
	if len(parts) != 3 {
		stdlog.ErrorF("parse reference string %s do not contain 2 slash", absRefStr)
		return ""
	}

	return util.StringConcat(parts[1], "/", parts[2])
}

type
// Deprecated
// migrate to: github.com/docker/docker/pkg/jsonmessage.JSONMessage
engineAPIv1dot43ImagesLoadResp struct {
	//{
	//	"Id": "sha256:abcdef123456...",
	//	"RepoTags": [
	//	"myimage:latest"
	//],
	//	"Message": ""
	//}
	Id          string   `json:"Id,omitempty"`       // what AI tells me is not-ok
	RepoTags    []string `json:"RepoTags,omitempty"` // what AI tells me is not-ok
	Message     string   `json:"Message,omitempty"`  // what AI tells me is not-ok
	Error       string   `json:"error,omitempty"`
	ErrorDetail any      `json:"errorDetail,omitempty"`
}
