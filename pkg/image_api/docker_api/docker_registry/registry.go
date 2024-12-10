package docker_registry

import "github.com/docker/docker/api/types/registry"

type Registry struct {
	authConfig registry.AuthConfig
	Addr       string   // registry address, domain:port, ip:port, domain, etc
	Auth       string   // base64 encoding, for HTTP API Header: registry.AuthHeader
	Project    []string // available projects
	Name       string   // identifier of current registry
}

func FromAuthConfig(authConfig registry.AuthConfig) (reg Registry) {
	reg.authConfig = authConfig
	reg.Addr = authConfig.ServerAddress
	auth, err := registry.EncodeAuthConfig(authConfig)
	if err != nil {
		return
	}
	reg.Auth = auth
	return
}

func (r *Registry) WithProject(project ...string) *Registry {
	r.Project = append(r.Project, project...)
	return r
}

func (r *Registry) WithName(name string) *Registry {
	r.Name = name
	switch name {
	case DockerHub:
		r.Project = append(r.Project, r.authConfig.Username)
	}
	return r
}

func (r *Registry) GetAuthConfig() registry.AuthConfig {
	return r.authConfig
}

const (
	DockerHub = "dockerhub"
	Harbor    = "harbor"
)
