package excel

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/xuri/excelize/v2"
)

type File struct {
	f *excelize.File
	w *writer
	r *reader
}

func (f *File) Done() *File {
	util.SilentCloseIO("excel ptr", f.f)
	return f
}

func (f *File) Err() error {
	switch {
	case f.w != nil:
		return f.w.First()

	case f.r != nil:
		return f.r.First()

	default:
		return nil
	}
}

func (f *File) DoF(do func(ef *excelize.File) error) *File {
	switch {
	case f.w != nil:
		f.w.Has(do(f.f))

	case f.r != nil:
		f.r.Has(do(f.f))

	default:
	}
	return f
}
