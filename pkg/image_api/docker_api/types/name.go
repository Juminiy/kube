package types

import (
	dockerdaemonnames "github.com/docker/docker/daemon/names"
	"regexp"
)

// regexp of artifact name: image_name:image_tag
const (
	// referred from ocispec: https://github.com/opencontainers/distribution-spec/blob/main/spec.md#workflow-categories
	RepositoryNameRegexp = "[a-z0-9]+((\\.|_|__|-+)[a-z0-9]+)*(\\/[a-z0-9]+((\\.|_|__|-+)[a-z0-9]+)*)*"
	ReferenceNameRegexp  = "[a-zA-Z0-9_][a-zA-Z0-9._-]{0,127}" // reference as tag, is a tag
	TagNameRegexp        = ReferenceNameRegexp                 // reference as tag, is a tag
	ArtifactNameRegexp   = RepositoryNameRegexp + ":" + TagNameRegexp
)

var (
	_repositoryNameRegexp = regexp.MustCompile(RepositoryNameRegexp)
	_referenceNameRegexp  = regexp.MustCompile(ReferenceNameRegexp)
	_artifactNameRegexp   = regexp.MustCompile(ArtifactNameRegexp)
)

// ValidArtifactName
// artifactName= image_name:image_tag
func ValidArtifactName(artifactName string) bool {
	return _artifactNameRegexp.MatchString(artifactName)
}

func ValidRepositoryName(repositoryName string) bool {
	return _repositoryNameRegexp.MatchString(repositoryName)
}

// ValidTagName
// referenceName= tagName
func ValidTagName(tagName string) bool {
	return _referenceNameRegexp.MatchString(tagName)
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
