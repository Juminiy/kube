package bufwriter

import (
	"github.com/Juminiy/kube/pkg/util/zerobuf"
	"github.com/spf13/cast"
	"strings"
)

type Writer struct {
	buf zerobuf.String
}

func New() *Writer {
	return &Writer{
		buf: zerobuf.Get(),
	}
}

func (w *Writer) String() string {
	defer w.buf.Free()
	return w.buf.String()
}

func (w *Writer) Line() *Writer {
	w.Byte('\n')
	return w
}

func (w *Writer) Space() *Writer {
	return w.Byte(' ')
}

func (w *Writer) Word(s ...string) *Writer {
	for i := range s {
		_, _ = w.buf.WriteString(s[i])
	}
	return w.Space()
}

func (w *Writer) Worda(v ...any) *Writer {
	for i := range v {
		_, _ = w.buf.WriteString(cast.ToString(v[i]))
	}
	return w.Space()
}

func (w *Writer) Words(s ...string) *Writer {
	for i := range s {
		_, _ = w.buf.WriteString(s[i])
		_ = w.Space()
	}
	return w.Space()
}

func (w *Writer) WordsSep(sep string, s ...string) *Writer {
	for i := range s {
		_, _ = w.buf.WriteString(s[i])
		if i == len(s)-1 && len(strings.TrimSpace(sep)) != 0 {
			break
		}
		_, _ = w.buf.WriteString(sep)
	}
	return w.Space()
}

func (w *Writer) Wordsa(v ...any) *Writer {
	for i := range v {
		_, _ = w.buf.WriteString(cast.ToString(v[i]))
		_ = w.Space()
	}
	return w.Space()
}

func (w *Writer) WordsaSep(sep string, v ...any) *Writer {
	for i := range v {
		_, _ = w.buf.WriteString(cast.ToString(v[i]))
		if i == len(v)-1 && len(strings.TrimSpace(sep)) != 0 {
			break
		}
		_, _ = w.buf.WriteString(sep)
	}
	return w.Space()
}

func (w *Writer) Byte(b byte) *Writer {
	_ = w.buf.WriteByte(b)
	return w
}
