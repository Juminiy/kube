package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"strings"
)

/*
sudo docker tag jammy-env:v1.6 10.112.121.243:8111/library/jammy-env:v1.6
sudo docker push 10.112.121.243:8111/library/jammy-env:v1.6
*/

/*
Address by tag: [loginServerUrl]/[repository][:tag]
Address by digest: [loginServerUrl]/[repository@sha256][:digest]
*/

type Artifact struct {
	// ex. docker.io
	Registry string

	// ex. Dockerhub is username, Harbor is project
	Project string

	// must use Artifact.SetName
	Name string
	// alias of Name
	repository string

	// must use Artifact.SetTag
	Tag string

	// must use Artifact.SetDigest
	Digest string
	// alias of Digest
	sha256Digest string

	// concat of `${Name}:${Tag}` or `${Name}:${Digest}`
	artifact string
}

// ParseToArtifact
// +example: jammy-env
// +example: jammy-env:v1.6
// +example: library/jammy-env:v1.6
// +example: 10.112.121.243:8111/library/jammy-env:v1.6
// +example: 10.112.121.243:8111/library/jammy-env@sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966
func ParseToArtifact(s string) (arti Artifact) {
	artiParts := strings.Split(s, "/")
	switch len(artiParts) {
	case 1:
		arti.parseArtifact(artiParts[0])

	case 2:
		arti.SetProject(artiParts[0]).parseArtifact(artiParts[1])

	case 3:
		arti.SetRegistry(artiParts[0]).SetProject(artiParts[1]).parseArtifact(artiParts[2])

	default:
		// invalid artifact, neither refStr nor digestStr
	}
	return
}

func (a *Artifact) parseArtifact(s string) {
	byColon := strings.Split(s, ":")
	byAt := strings.Split(s, "@")
	if len(byColon) == 1 && len(byAt) == 1 {
		a.SetName(s).SetTag(TagLatest)
	} else if len(byAt) == 2 {
		a.setArtifact(s).SetName(byAt[0]).SetDigest(byAt[1])
	} else if len(byColon) == 2 {
		if byColon[0] == "sha256" {
			a.SetDigest(s)
		} else {
			a.setArtifact(s).SetName(byColon[0]).SetTag(byColon[1])
		}
	}
}

func (a *Artifact) AbsRefStr() string {
	if len(a.Registry) > 0 {
		return a.Registry + "/" + a.RefStr()
	}
	return a.RefStr()
}

func (a *Artifact) RefStr() string {
	if len(a.Project) > 0 {
		return a.Project + "/" + a.artifact
	}
	return a.artifact
}

func (a *Artifact) AbsDigestStr() string {
	if len(a.Registry) > 0 {
		return a.Registry + "/" + a.DigestStr()
	}
	return a.DigestStr()
}

func (a *Artifact) DigestStr() string {
	if len(a.Project) > 0 && len(a.repository) > 0 {
		return a.Project + "/" + a.repository + "@" + a.sha256Digest
	} else if len(a.repository) > 0 {
		return a.repository + "@" + a.sha256Digest
	}
	return a.sha256Digest
}

// +example: http://10.112.121.243:8111
// +example: 10.112.121.243:8111/
// +example: 10.112.121.243:8111
func (a *Artifact) SetRegistry(s string) *Artifact {
	a.Registry = trimSSlash(util.TrimProto(s))
	return a
}

// +example: /library/
// +example: library
func (a *Artifact) SetProject(s string) *Artifact {
	a.Project = trimSlash(s)
	return a
}

// +example: /jammy-env
// +example: jammy-env
func (a *Artifact) SetName(s string) *Artifact {
	s = trimPSlash(s)
	a.repository = s
	a.Name = s
	return a
}

// +example: :v1.6
// +example: v1.6
func (a *Artifact) SetTag(s string) *Artifact {
	s = strings.TrimPrefix(s, ":")
	a.Tag = s
	if len(a.repository) > 0 && len(s) > 0 {
		a.setArtifact(a.repository + ":" + s)
	}
	return a
}

// +example: @sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966
// +example: sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966
func (a *Artifact) SetDigest(s string) *Artifact {
	s = strings.TrimPrefix(s, "@")
	a.sha256Digest = s
	a.Digest = s
	return a
}

func (a *Artifact) SetArtifact(s string) *Artifact {
	a.parseArtifact(s)
	return a
}

// +example: /jammy-env:v1.6
// +example: jammy-env:v1.6
func (a *Artifact) setArtifact(s string) *Artifact {
	a.artifact = trimPSlash(s)
	return a
}

func trimPSlash(s string) string {
	return strings.TrimPrefix(s, "/")
}

func trimSSlash(s string) string {
	return strings.TrimSuffix(s, "/")
}

func trimSlash(s string) string {
	return trimPSlash(trimSSlash(s))
}

const (
	TagLatest = "latest"
)
