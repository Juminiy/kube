package docker_api

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_client"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	"io"
)

type ImportImageResp struct {
	RequestRefStr string
	LoadedRefStr  string
	// +disable in Client.ImportImageV2
	PushStatus string
	PushDigest string
	Inspect    types.ImageInspect
	PushResp   PushImageOfficialAPIResp
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
		if len(inspect.RepoTags) > 0 {
			loadedImageID = inspect.RepoTags[0]
		}
		for i := range inspect.RepoTags {
			if inspect.RepoTags[i] != absRefStr {
				loadedRefStr = inspect.RepoTags[i]
				break
			}
		}
		resp.Inspect = inspect
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
		if len(inspect.RepoTags) > 0 {
			loadedImageID = inspect.RepoTags[0]
		}
		for i := range inspect.RepoTags {
			if inspect.RepoTags[i] != absRefStr {
				loadedRefStr = inspect.RepoTags[i]
				break
			}
		}
		resp.Inspect = inspect
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
	resp.PushDigest = pushImageResp.GetDigest()
	resp.PushResp = pushImageResp
	return
}

type ImportImageRespV3 struct {
	RequestRefStr string
	LoadedRefStr  string
	Inspect       types.ImageInspect
	PushResp      docker_client.ImagePushResp
}

func (c *Client) ImportImageV3(absRefStr string, input io.Reader) (resp ImportImageRespV3, err error) {
	resp.RequestRefStr = absRefStr
	loadResp, err := c.apiClient.ImageLoad(input)
	if err != nil {
		return
	}
	imageLoaded := loadResp.GetImageLoad()
	resp.Inspect, err = c.InspectImage(imageLoaded)
	if err != nil {
		return
	}
	loadRefStr := loadedRefStr(&resp.Inspect)
	if len(loadRefStr) == 0 {
		return resp, errImageNotFound
	}
	resp.LoadedRefStr = loadRefStr
	pushResp, err := c.CreateImageTagV3(absRefStr, resp.LoadedRefStr)
	if err != nil {
		return
	}
	resp.PushResp = pushResp.GetImagePushResp()
	return
}

func loadedRefStr(inspect *types.ImageInspect) string {
	if len(inspect.RepoTags) == 0 {
		return ""
	}
	return inspect.RepoTags[0]
}
