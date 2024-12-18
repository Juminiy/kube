package docker_client

import (
	"context"
	"github.com/docker/docker/api/types"
	"io"
)

type APIClient interface {
	EventAPIClient
}

type EventAPIClient interface {
	ImageAPIClient
}

type ImageAPIClient interface {
	ImageLoad(input io.Reader) (EventResp, error)
	ImagePush(refStr string) (EventResp, error)
	ImageCreate(parentRefStr string) (EventResp, error)
	ImagePull(refStr string) (EventResp, error)
	ImageBuild(input io.Reader, options types.ImageBuildOptions, ctx ...context.Context) (EventResp, error)
}
