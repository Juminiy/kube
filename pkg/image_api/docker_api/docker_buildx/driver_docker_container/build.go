package driver_docker_container

type BuildImage interface {
	RunBuild() error
}
