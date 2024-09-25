package stdlog

import (
	"github.com/docker/docker/pkg/jsonmessage"
	"os"
)

type stream struct{}

func (s *stream) Write(bs []byte) (int, error) {
	return os.Stdout.Write(bs)
}

func (s *stream) FD() uintptr {
	return os.Stdout.Fd()
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
