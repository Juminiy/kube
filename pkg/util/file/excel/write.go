package excel

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/xuri/excelize/v2"
	"sync"
)

type writer struct {
	sheetName sync.Map
	*util.ErrHandle
}

func (w *writer) hasSheet(sheetName string) bool {
	_, ok := w.sheetName.LoadOrStore(sheetName, util.NilStruct())
	return ok
}

const (
	defaultSheetName = "Sheet1"
)

func NewWriter() *File {
	return &File{
		f: excelize.NewFile(),
		w: &writer{
			sheetName: sync.Map{},
			ErrHandle: util.NewErrHandle(),
		},
	}
}

func (f *File) DeleteDefaultSheet() *File {
	f.w.Has(f.f.DeleteSheet(defaultSheetName))
	return f
}

func (f *File) AppendRow(sheetName string, row ...[]any) *File {
	if f.w.Has() {
		return f
	}
	if !f.w.hasSheet(sheetName) {
		_, err := f.f.NewSheet(sheetName)
		if f.w.Has(err) {
			return f
		}
	}
	for i := range row {
		cellName, err := excelize.CoordinatesToCellName(1, i+1)
		if f.w.Has(err,
			f.f.SetSheetRow(sheetName, cellName, &row[i])) {
			return f
		}
	}
	return f
}
