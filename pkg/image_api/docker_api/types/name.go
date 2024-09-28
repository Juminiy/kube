package types

import (
	dockerdaemonnames "github.com/docker/docker/daemon/names"
	"regexp"
)

// regexp of artifact name: image_name:image_tag
const (
	ArtifactNameRegexp = ""
)

var (
	_artifactNameRegexp = regexp.MustCompile("")
)

// ValidArtifactName
// artifactName= image_name:image_tag
func ValidArtifactName(artifactName string) bool {
	return _artifactNameRegexp.MatchString(artifactName)
}

// regexp of container name
const (
	// ContainerNameRegexp
	// referred from: github.com/docker/docker/daemon/names.RestrictedNameChars
	ContainerNameRegexp = dockerdaemonnames.RestrictedNameChars
)

var (
	_containerNameRegexp = dockerdaemonnames.RestrictedNamePattern
)

func ValidContainerName(containerName string) bool {
	return _containerNameRegexp.MatchString(containerName)
}

func ValidVolumeName(volumeName string) bool {
	return _containerNameRegexp.MatchString(volumeName)
}
