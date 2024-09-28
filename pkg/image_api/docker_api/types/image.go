package types

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	distributionref "github.com/distribution/reference"
	dockerimage "github.com/docker/docker/image"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"strings"
)

type ImageRef struct {
	Registry   string
	Project    string
	Repository string
	Tag        string

	Platform *ocispec.Platform
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

func (r *ImageRef) Rep() string {
	return r.String()
}

type ImageRep interface {
	dockerimage.ID
	distributionref.Named
	AvaRep
}
