package excel

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/xuri/excelize/v2"
	"os"
	testing2 "testing"
)

func TestNewWriter(t *testing2.T) {
	err := NewWriter().AppendRow("sheet-test",
		[]any{"Brand", "Sale", "Price"},
		[]any{"Apple", 10, 6.66},
		[]any{"Xiaomi", 9, 3.33},
		[]any{"Huawei", 5, 8.88}).
		DeleteDefaultSheet().
		DoF(func(ef *excelize.File) error {
			return ef.SaveAs("testdata/sheet1.xlsx")
		}).Err()
	util.Must(err)
}

func TestNewReader(t *testing2.T) {
	fptr, err := os.Open("testdata/sheet1.xlsx")
	util.Must(err)
	defer util.SilentCloseIO("file ptr", fptr)
	var rows [][]string
	var allrows [][][]string
	err = NewReader(fptr).AllSheetRows(&allrows).SheetRows("sheet-test", &rows).Err()
	util.Must(err)
	t.Log(allrows)
	t.Log(rows)
}
