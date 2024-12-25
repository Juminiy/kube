package kinternal

import "io"

type StringBuilder interface {
	Build() string
}

type WriteBuilder interface {
	Build(w io.Writer)
}
