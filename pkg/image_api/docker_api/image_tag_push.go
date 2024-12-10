package docker_api

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_client"
	"io"
)

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

func (c *Client) CreateImageTagV3(toAbsRefStr, fromAbsRefStr string) (resp docker_client.EventResp, err error) {
	err = c.cli.ImageTag(c.ctx, fromAbsRefStr, toAbsRefStr)
	if err != nil {
		return
	}
	return c.apiClient.ImagePush(toAbsRefStr)
}

type TagPushImageResp struct {
	Artifact
	docker_client.ImagePushResp
}

func (c *Client) tagImageFromRefStr(refStr string) (tagPushResp TagPushImageResp, err error) {
	arti := ParseToArtifact(refStr)
	if len(arti.Registry) == 0 {
		arti.SetRegistry(c.reg.Addr)
	}
	if len(arti.Project) == 0 && len(c.reg.Project) > 0 {
		arti.SetProject(c.reg.Project[0])
	}
	tagPushResp.Artifact = arti
	tagPush, err := c.CreateImageTagV3(arti.AbsRefStr(), refStr)
	if err != nil {
		return
	}
	tagPushResp.ImagePushResp = tagPush.GetImagePushResp()
	return
}
