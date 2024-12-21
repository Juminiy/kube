package ksql

type B interface {
	Build() string
}

type builder struct{}

func (b builder) Build() string {
	return ""
}
