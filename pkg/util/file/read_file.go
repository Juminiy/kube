package file

import (
	"bufio"
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/util"
	"io"
	"os"
)

type Reader struct {
	filePath string
	fptr     *os.File
	rd       *bufio.Reader
}

func NewReader(filePath string) *Reader {
	if len(filePath) == 0 ||
		internal_api.FileNotExist(filePath) {
		return nil
	}
	fptr, err := os.Open(filePath)
	util.Must(err)
	return &Reader{
		filePath: filePath,
		fptr:     fptr,
		rd:       bufio.NewReader(fptr),
	}
}

func (r *Reader) Done() {
	util.SilentCloseIO("file ptr", r.fptr)
}

func (r *Reader) NextLine2() (lineStr string, hasNext bool) {
	lineStr, err := r.rd.ReadString('\n')
	return lineStr, r.doError(err)
}

func (r *Reader) NextLine(s *string) bool {
	lineStr, hasNext := r.NextLine2()
	*s = lineStr
	return hasNext
}

func (r *Reader) NextRune2() (un rune, hasNext bool) {
	un, _, err := r.rd.ReadRune()
	return un, r.doError(err)
}

func (r *Reader) NextRune(un *rune) bool {
	runer, hasNext := r.NextRune2()
	*un = runer
	return hasNext
}

func (r *Reader) NextByte2() (b byte, hasNext bool) {
	b, err := r.rd.ReadByte()
	return b, r.doError(err)
}

func (r *Reader) NextByte(b *byte) bool {
	bb, hasNext := r.NextByte2()
	*b = bb
	return hasNext
}

func (r *Reader) doError(err error) (hasNext bool) {
	switch {
	case err == nil:
		hasNext = true

	case err == io.EOF:
		hasNext = false

	default:
		util.Must(err)
	}
	return
}
