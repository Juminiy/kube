package docker_file

import "io"

// docker_file for reinterpret Dockerfile to human desc

type Builder interface {
	Build(w io.StringWriter)
}
