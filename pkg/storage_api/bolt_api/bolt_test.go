package bolt_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/boltdb/bolt"
	"testing"
)

var Enc = safe_json.Pretty
var B2S = util.Bytes2StringNoCopy
var S2B = util.String2BytesNoCopy

func TestBolt(t *testing.T) {
	cli, err := New("bolt.db")
	util.Must(err)
	defer util.SilentCloseIO("bolt db", cli)
	t.Log(Enc(cli.Info()))
	err = cli.Update(func(tx *bolt.Tx) error {
		btk, err := tx.CreateBucketIfNotExists(S2B("rr"))
		if err != nil {
			return err
		}
		err = btk.Put(S2B("topic-deepseek"), S2B("{\"json_str\":\"js\"}"))
		if err != nil {
			return err
		}

		return nil
	})
	util.Must(err)

	var noneV, deepseekV []byte
	err = cli.View(func(tx *bolt.Tx) error {
		btk := tx.Bucket(S2B("rr"))
		noneV = btk.Get(S2B("topic-none"))
		deepseekV = btk.Get(S2B("topic-deepseek"))
		return nil
	})
	t.Log(B2S(noneV))
	t.Log(B2S(deepseekV))
}
