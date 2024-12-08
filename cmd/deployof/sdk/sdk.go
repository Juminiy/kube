package sdk

import "github.com/Juminiy/kube/pkg/image_api/harbor_api"

var registryProject string

func Init(project string) {
	registryProject = project
}

func GetImageAddr(name, tag string) string {
	return harbor_api.ArtifactURI{
		Project:    registryProject,
		Repository: name,
		Tag:        tag,
	}.String()
}
