package util

import "context"

var (
	_todoContext       = context.TODO()
	_backgroundContext = context.Background()
)

func TODOContext() context.Context {
	return _todoContext
}

func BackgroundContext() context.Context {
	return _backgroundContext
}
