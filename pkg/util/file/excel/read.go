package excel

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/xuri/excelize/v2"
	"io"
	"sync"
)

type reader struct {
	rows sync.Map
	*util.ErrHandle
}

func NewReader(rd io.Reader) *File {
	f := &File{r: &reader{
		rows:      sync.Map{},
		ErrHandle: util.NewErrHandle(),
	}}
	ff, ferr := excelize.OpenReader(rd)
	f.f = ff
	f.r.Has(ferr)
	return f
}

var (
	errResultParamNil = errors.New("result param is nil")
)

func (f *File) AllSheetRows(allRows *[][][]string) *File {
	if f.r.Has() {
		return f
	}
	if allRows == nil {
		f.r.Has(errResultParamNil)
		return f
	}
	list := f.f.GetSheetList()
	for i := range list {
		var rows [][]string
		f.SheetRows(list[i], &rows)
		*allRows = append(*allRows, rows)
	}
	return f
}

// SheetRows
// +param sheetName
// +result sheetRows
func (f *File) SheetRows(sheetName string, sheetRows *[][]string) *File {
	if f.r.Has() {
		return f
	}
	if sheetRows == nil {
		f.r.Has(errResultParamNil)
		return f
	}
	cacheRows, ok := f.r.rows.Load(sheetName)
	if ok {
		*sheetRows = cacheRows.([][]string)
		return f
	}
	readRows, err := f.f.GetRows(sheetName)
	if f.r.Has(err) {
		return f
	}
	*sheetRows = readRows
	f.r.rows.Store(sheetName, readRows)
	return f
}
