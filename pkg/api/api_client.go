package api

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
)

type ClientInterface interface {
	WithContext(context.Context) ClientInterface
	GCable
}

type LevelLogger interface {
	Debug(...any)
	Info(...any)
	Warn(...any)
	Error(...any)
	Fatal(...any)
	Panic(...any)
}

type GCable interface {
	GC(...util.Func)
}
