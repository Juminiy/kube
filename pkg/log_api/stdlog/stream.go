package stdlog

import (
	"github.com/docker/docker/pkg/jsonmessage"
)

// conform and consistency with Docker pkg jsonmessage
type stream struct{}

func (s *stream) Write(bs []byte) (int, error) {
	return _logDefaultOut.Write(bs)
}

func (s *stream) FD() uintptr {
	return _logDefaultOut.Fd()
}

func (s *stream) IsTerminal() bool {
	return true
}

var (
	_stream = &stream{}
)

func Stream() jsonmessage.Stream {
	return _stream
}
