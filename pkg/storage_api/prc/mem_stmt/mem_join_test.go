package mem_stmt

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

var usrOrderHdr = []string{
	"user.id", "user.name", "order.id", "order.user_id", "order.product",
}

var tblUser = [][]string{
	{"1", "Bob"},
	{"2", "Anna"},
	{"3", "Sam"},
}

var tblOrder = [][]string{
	{"1", "1", "Apple"},
	{"2", "2", "Orange"},
	{"3", "4", "Banana"},
	{"4", "8", "Juice"},
}

func d2table(t *testing.T, hdr []string, rows [][]string) {
	t.Log(hdr)
	for _, ri := range rows {
		t.Log(ri, '\n')
	}
	util.TestLongHorizontalLine(t)
}

func TestJoin(t *testing.T) {
	j := Join{}
	d2table(t, usrOrderHdr, j.Cross(tblUser, tblOrder))
	d2table(t, usrOrderHdr, j.Inner(tblUser, tblOrder, 0, 1))
	d2table(t, usrOrderHdr, j.Left(tblUser, tblOrder, 0, 1))
	d2table(t, usrOrderHdr, j.Right(tblOrder, tblUser, 1, 0))
}
