package korm

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

type K struct {
	ID     uint    `k:"rowid"`
	Name   string  `k:"name"`
	Desc   string  `k:"desc"`
	Rate   float64 `k:"rate"`
	Stat   int     `k:"status"`
	Extras string  `k:"extras"`
}

func (K) TName() string { return `tbl_k` }

func TestOpen(t *testing.T) {
	tx, err := Open()
	util.Must(err)
	defer util.SilentCloseIO("db", tx)
	err = tx.Create(&K{
		Name:   "Neko",
		Desc:   "FavMiaomiao",
		Rate:   0.99,
		Stat:   1,
		Extras: `{"platform": ["Bilibili", "SnapChat", "Pinterest"]}`,
	})
	util.Must(err)
}
