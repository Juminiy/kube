package docker_file

import (
	"bytes"
	"io"
)

// docker_file for reinterpret Dockerfile to human desc

type Builder interface {
	Build(w io.StringWriter)
}

type Doc struct{}

func (d *Doc) Reader(bs ...byte) *bytes.Reader {
	return bytes.NewReader(bs)
}
