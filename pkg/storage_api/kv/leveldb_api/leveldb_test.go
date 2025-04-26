package leveldb_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/syndtr/goleveldb/leveldb"
	leveldbutil "github.com/syndtr/goleveldb/leveldb/util"
	"strconv"
	"testing"
)

var _S2B = util.String2BytesNoCopy
var _B2S = util.Bytes2StringNoCopy

var _db *leveldb.DB

func init() {
	db, err := New("ldb.db")
	util.Must(err)
	_db = db
}

func TestLeveldb(t *testing.T) {
	defer util.SilentCloseIO("level db file", _db)
	//util.Must(db.Put(_S2B("lkey"), _S2B("lval"), nil))
	util.Must(_db.Delete(_S2B("lkey"), nil))
	lval, err := _db.Get(_S2B("lkey"), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			t.Logf("key[%s] not found", "lkey")
		}
	}
	t.Log(_B2S(lval))
}

func TestBatch(t *testing.T) {
	defer util.SilentCloseIO("level db file", _db)
	ltx, err := _db.OpenTransaction()
	util.Must(err)
	for i := range 114514 {
		th := strconv.Itoa(i)
		util.Must(ltx.Put(_S2B("lkey_"+th), _S2B("lval_"+th), nil))
	}
	util.Must(ltx.Commit())
	//_db.Write(&leveldb.Batch{}, nil)
}

func TestIterRange(t *testing.T) {
	defer util.SilentCloseIO("level db file", _db)
	iterRg := _db.NewIterator(&leveldbutil.Range{
		Start: _S2B("lkey_70000"),
		Limit: _S2B("lkey_79999"),
	}, nil)
	for iterRg.Next() {
		lkey, lval := iterRg.Key(), iterRg.Value()
		t.Logf("[%s]:[%s]", _B2S(lkey), _B2S(lval))
	}
}

func TestIterPrefix(t *testing.T) {
	defer util.SilentCloseIO("level db file", _db)
	iterPref := _db.NewIterator(leveldbutil.BytesPrefix(_S2B("lkey_808")), nil)
	for iterPref.Next() {
		lKey, lVal := iterPref.Key(), iterPref.Value()
		t.Logf("[%s]:[%s]", _B2S(lKey), _B2S(lVal))
	}
}
