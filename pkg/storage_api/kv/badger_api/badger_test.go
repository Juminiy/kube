package badger_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/dgraph-io/badger/v4"
	"testing"
)

var _S2B = util.String2BytesNoCopy
var _B2S = util.Bytes2StringNoCopy

func TestBadger(t *testing.T) {
	db, err := New("b.db")
	util.Must(err)
	defer util.SilentCloseIO("badger db file", db)
	//err = db.Update(func(txn *badger.Txn) error {
	//	return txn.Set(_S2B("keyd"), _S2B("vald"))
	//})
	//util.Must(err)
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Delete(_S2B("keyd"))
	})
	util.Must(err)
	var bv []byte
	err = db.View(func(txn *badger.Txn) error {
		bitem, err := txn.Get(_S2B("keyd"))
		if err != nil {
			return err
		}
		bv, err = bitem.ValueCopy(bv)
		return err
	})
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			t.Logf("key[%s]: not found", "keyd")
		}
	}
	t.Log(_B2S(bv))
}
