package file

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/util"
	"os"
)

type Writer struct {
	filePath string
	fptr     *os.File
}

func (w *Writer) Done() {
	util.SilentCloseIO("file ptr", w.fptr)
}

func (w *Writer) Line() *Writer {
	w.bytes('\n')
	return w
}

func (w *Writer) Word(s string) *Writer {
	return w.words(s)
}

func (w *Writer) Words(s ...string) *Writer {
	return w.words(s...)
}

func (w *Writer) bytes(bs ...byte) *Writer {
	_, err := w.fptr.Write(bs)
	util.Must(err)
	return w
}

func (w *Writer) words(s ...string) *Writer {
	for i := range s {
		_, err := w.fptr.WriteString(s[i])
		util.Must(err)
		w.bytes(' ')
	}
	return w
}

func CodeGen(filePath string) *Writer {
	if len(filePath) == 0 {
		return nil
	}
	fptr, err := internal_api.OverwriteCreateFile(filePath)
	util.Must(err)
	w := &Writer{
		filePath: filePath,
		fptr:     fptr,
	}

	return w
}
