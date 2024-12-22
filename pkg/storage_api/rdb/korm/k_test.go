package korm

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

type K struct {
	ID     uint    `k:"col:rowid" json:"id"`
	Name   string  `k:"col:name" json:"name"`
	Desc   string  `k:"col:desc" json:"desc"`
	Rate   float64 `k:"col:rate" json:"rate"`
	Stat   int     `k:"col:status" json:"status"`
	Extras string  `k:"col:extras" json:"extras"`
}

func (K) TName() string { return `tbl_k` }

var _tx I

func init() {
	tx, err := Open()
	util.Must(err)
	_tx = tx
}

func TestSession_Create(t *testing.T) {
	nekoFav := K{
		Name:   "Neko",
		Desc:   "FavMiaomiao",
		Rate:   0.93,
		Stat:   1,
		Extras: `{"platform": ["Bilibili", "SnapChat", "Pinterest"]}`,
	}
	err := _tx.Create(&nekoFav)
	util.Must(err)
	t.Log(safe_json.Pretty(nekoFav))
}

func TestSession_Create2(t *testing.T) {
	favList := []K{
		{
			Name:   "BraveShine",
			Desc:   "FavMusic",
			Rate:   0.78,
			Stat:   1,
			Extras: `{"platforms": ["Youtube", "Bilibili", "Spotify"]}`,
		},
		{
			Name:   "C++11",
			Desc:   "Technical Programming Lang",
			Rate:   0.75,
			Stat:   1,
			Extras: `{"platforms": ["Youtube", "Bilibili", "Stackoverflow"]}`,
		},
		{
			Name:   "LYY",
			Desc:   "FavFemale",
			Rate:   0.99,
			Stat:   1,
			Extras: `{"platforms": ["NEUQ", "2084team", "StudentCommitteeCadres"]}`,
		},
	}
	err := _tx.Create(&favList)
	util.Must(err)
	t.Log(safe_json.Pretty(favList))
}
