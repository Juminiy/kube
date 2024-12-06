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
	"github.com/docker/docker/errdefs"
	"io"
	"strings"
)

var (
	errImageNotFound = errors.New("image not found")
	errImageInternal = errors.New("image internal error")
)

type ExportImageResp struct {
	RequestRefStr      string
	NotFoundInRegistry bool          // absRefStr not found in provided registry.
	NotFoundInLocal    bool          // absRefStr not found in docker-server local.
	ImagePulledInfo    string        // image pulled information from provided registry.
	ImageFileReader    io.ReadCloser // compressed (.tar, .tar.gz) file reader, It's up to the caller to handle the io.ReadCloser and close it properly.
}

type errStrParser struct{}

var _errStrParser errStrParser

// +example:
// "Error response from daemon: unknown: artifact library/hello:v1.0 not found"
func (errStrParser) ArtifactNotFound(err error) bool {
	if err != nil {
		words := strings.Split(err.Error(), " ")
		indexArti, indexNot, indexFound := -1, -1, -1
		for i := range words {
			switch words[i] {
			case "artifact":
				indexArti = i
			case "not":
				indexNot = i
			case "found":
				indexFound = i
			}
		}
		return indexArti > 0 && indexNot > 0 && indexFound > 0 && indexFound == indexNot+1
	}
	return false
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
			resp.NotFoundInRegistry = _errStrParser.ArtifactNotFound(err)
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

type ImportImageResp struct {
	RequestRefStr string
	LoadedRefStr  string
	PushStatus    string
	Digests       []string
	Digest        string
}

// ImportImage
// +param input
//
// (1). make sure the input file(.tar.gz) metadata attribute image name format is: registry/project/image_name:image_tag
//
//	(1.1). or will create new image tag to local
//
// (2). push image to registry
func (c *Client) ImportImage(absRefStr string, input io.Reader) (resp ImportImageResp, err error) {
	resp.RequestRefStr = absRefStr
	loadResp, err := c.cli.ImageLoad(c.ctx, input, false)
	if err != nil {
		return resp, err
	}
	defer util.SilentCloseIO("docker image load", loadResp.Body)

	var loadedRefStr string
	if loadResp.Body != nil && loadResp.JSON {
		// return json message
		//stdlog.Debug("docker image loadResp format: json")
		// +example1
		// resp.body.body.body.src.r.buf JSON{"stream":"Loaded image ID: sha256:d2c94e258dcb3c5ac2798d32e1249e42ef01cba4841c2234249495f87264ac5a\n"}
		// STDOUT: Loaded image ID: sha256:249f59e1dec7f7eacbeba4bb9215b8000e4bdbb672af523b3dacc89915b026ae
		// +example2
		// resp.body.body.body.src.r.buf JSON{"errorDetail":{"message":"open /var/lib/docker/tmp/docker-import-3823099216/repositories: no such file or directory"},"error":"open /var/lib/docker/tmp/docker-import-3823099216/repositories: no such file or directory"}
		/*ignoreJSONMessageErr := jsonmessage.DisplayJSONMessagesToStream(loadResp.Body, stdlog.Stream(), nil)
		if ignoreJSONMessageErr != nil {
			stdlog.Warn(ignoreJSONMessageErr)
		}*/
		loadedImageID := docker_internal.GetImageIDFromImageLoadResp(loadResp.Body)
		inspect, err := c.InspectImage(loadedImageID)
		if err != nil {
			return resp, err
		}
		for i := range inspect.RepoTags {
			if inspect.RepoTags[i] != absRefStr {
				loadedRefStr = inspect.RepoTags[i]
				break
			}
		}
		resp.Digests = inspect.RepoDigests
	} else {
		//stdlog.Warn("docker client API Version too old, can not create tag furthermore")
		//stdlog.Debug("docker image loadResp format: plain text")
		_, err = io.Copy(stdlog.Stream(), loadResp.Body)
		if err != nil {
			return resp, err
		}
	}
	if len(loadedRefStr) == 0 {
		return resp, errImageInternal
	}
	resp.LoadedRefStr = loadedRefStr

	pushImageRc, err := c.CreateImageTag(absRefStr, loadedRefStr)
	defer util.SilentCloseIO("pushImageReadCloser", pushImageRc)
	if err != nil {
		return resp, err
	}
	resp.PushStatus = docker_internal.GetStatusFromImagePushResp(pushImageRc)
	return
}

func (c *Client) ImportImageV2(absRefStr string, input io.Reader) (resp ImportImageResp, err error) {
	resp.RequestRefStr = absRefStr
	loadResp, err := c.cli.ImageLoad(c.ctx, input, false)
	if err != nil {
		return resp, err
	}
	defer util.SilentCloseIO("docker image load", loadResp.Body)

	var loadedRefStr string
	if loadResp.Body != nil && loadResp.JSON {
		// return json message
		//stdlog.Debug("docker image loadResp format: json")
		// +example1
		// resp.body.body.body.src.r.buf JSON{"stream":"Loaded image ID: sha256:d2c94e258dcb3c5ac2798d32e1249e42ef01cba4841c2234249495f87264ac5a\n"}
		// STDOUT: Loaded image ID: sha256:249f59e1dec7f7eacbeba4bb9215b8000e4bdbb672af523b3dacc89915b026ae
		// +example2
		// resp.body.body.body.src.r.buf JSON{"errorDetail":{"message":"open /var/lib/docker/tmp/docker-import-3823099216/repositories: no such file or directory"},"error":"open /var/lib/docker/tmp/docker-import-3823099216/repositories: no such file or directory"}
		/*ignoreJSONMessageErr := jsonmessage.DisplayJSONMessagesToStream(loadResp.Body, stdlog.Stream(), nil)
		if ignoreJSONMessageErr != nil {
			stdlog.Warn(ignoreJSONMessageErr)
		}*/
		loadedImageID := docker_internal.GetImageIDFromImageLoadResp(loadResp.Body)
		inspect, err := c.InspectImage(loadedImageID)
		if err != nil {
			return resp, err
		}
		for i := range inspect.RepoTags {
			if inspect.RepoTags[i] != absRefStr {
				loadedRefStr = inspect.RepoTags[i]
				break
			}
		}
		resp.Digests = inspect.RepoDigests
	} else {
		//stdlog.Warn("docker client API Version too old, can not create tag furthermore")
		//stdlog.Debug("docker image loadResp format: plain text")
		_, err = io.Copy(stdlog.Stream(), loadResp.Body)
		if err != nil {
			return resp, err
		}
	}
	if len(loadedRefStr) == 0 {
		return resp, errImageInternal
	}
	resp.LoadedRefStr = loadedRefStr

	pushImageResp, err := c.CreateImageTagV2(absRefStr, loadedRefStr)
	if err != nil {
		return resp, err
	}
	resp.Digest = pushImageResp.GetDigest()
	return
}

func (c *Client) CreateImageTag(toAbsRefStr, fromAbsRefStr string) (io.ReadCloser, error) {
	err := c.cli.ImageTag(c.ctx, fromAbsRefStr, toAbsRefStr)
	if err != nil {
		return nil, err
	}
	return c.pushImage(toAbsRefStr)
}

func (c *Client) CreateImageTagV2(toAbsRefStr, fromAbsRefStr string) (resp PushImageOfficialAPIResp, err error) {
	err = c.cli.ImageTag(c.ctx, fromAbsRefStr, toAbsRefStr)
	if err != nil {
		return resp, err
	}
	return c.pushImageV2(toAbsRefStr)
}

// InspectImage
// +param imageID or imageName
func (c *Client) InspectImage(imageID string) (types.ImageInspect, error) {
	// discard raw []byte, no usage anymore
	imageInspect, _, err := c.cli.ImageInspectWithRaw(c.ctx, imageID)
	return imageInspect, err
}

type HostImageGCFunc util.Func

// HostImageGC
// cli: docker rmi IMAGE_ID
// maybe quota by:
//
// 1. image CREATED: since, before
// 2. image SIZE: bytes(B)
// 3. cache algorithm policy: lru, lfu
// 4. host disk: bytes(B)
func (c *Client) HostImageStorageGC(gcFunc ...HostImageGCFunc) {

}

func (c *Client) pullImage(absRefStr string) (io.ReadCloser, error) {
	return c.cli.ImagePull(c.ctx, absRefStr, image.PullOptions{
		All:          false,
		RegistryAuth: c.xRegistryAuth,
	})
}

func (c *Client) pushImage(absRefStr string) (io.ReadCloser, error) {
	return c.cli.ImagePush(c.ctx, absRefStr, image.PushOptions{
		All:          false,
		RegistryAuth: c.xRegistryAuth,
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

// in registry: search images by filter
func (c *Client) searchImageByRef(absRefStr string) ([]registry.SearchResult, error) {
	return c.cli.ImageSearch(c.ctx, "", registry.SearchOptions{
		RegistryAuth:  c.xRegistryAuth,
		PrivilegeFunc: nil,
		Filters:       getRefFilter(absRefStr),
	})
}

func getRefFilter(absRefStr string) filters.Args {
	if !validAbsRefStr(absRefStr) {
		return filters.NewArgs(filters.KeyValuePair{
			Key:   docker_internal.FilterReference,
			Value: docker_internal.ReferenceNone,
		})
	}
	return filters.NewArgs(filters.KeyValuePair{
		Key:   docker_internal.FilterReference,
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
)
