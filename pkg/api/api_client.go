// Package api: Unify the API configuration and API call methods, and hide the internal details of the API
package api

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
)

type Client interface {
	WithContext(context.Context) Client
	GCable
}

type GCable interface {
	GC(...util.Func)
}

type LevelLogger interface {
	Debug(...any)
	Info(...any)
	Warn(...any)
	Error(...any)
	Fatal(...any)
	Panic(...any)
}

type InternalLogger interface {
	Info(...any)
	Infof(string, ...any)
	Error(...any)
	Errorf(string, ...any)
}

type InternalLoggerV2 interface {
	Info(...any)
	InfoF(string, ...any)
	Error(...any)
	ErrorF(string, ...any)
}

func New(...any) (*Client, error) {
	return nil, nil
}
