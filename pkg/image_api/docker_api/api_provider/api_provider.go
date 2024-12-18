package api_provider

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_client"
	dockercli "github.com/docker/docker/client"
)

type APIProvider struct {
	// official SDK APIProvider
	SDK dockercli.APIClient

	// official API APIProvider, alias of unofficial SDK
	API docker_client.APIClient
}
