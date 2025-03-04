package gorm_api

import (
	"database/sql"
	"encoding/json"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
	"time"
)

type Name string

func (n Name) MarshalJSON() ([]byte, error) {
	if n == "" || n == "null" {
		return nil, nil
	}
	return util.String2BytesNoCopy(string(n)), nil
}

func (n *Name) UnmarshalJSON(b []byte) error {
	if str := util.Bytes2StringNoCopy(b); str != "null" {
		*n = Name(str)
	}
	return nil
}

func TestMagicType(t *testing.T) {
	n := Name("null")
	bs, err := json.Marshal(n)
	Err(t, err)
	t.Logf("%s", bs)
}

func TestTimeAlias(t *testing.T) {
	vList := []string{
		"null",
		Enc(struct { // not ok
			ID   uint
			Time sql.NullTime
		}{ID: 10,
			Time: sql.NullTime{Time: time.Now()}}),

		Enc(struct {
			ID   uint
			Time time.Time
		}{ID: 10,
			Time: time.Now()}),

		Enc(struct {
			ID   uint
			Time int64
		}{ID: 10,
			Time: time.Now().Unix(),
		}),

		Enc(struct { // not ok
			ID   uint
			Time string
		}{ID: 10,
			Time: time.Now().String(),
		}),
	}
	for i, timeRep := range vList {
		var v struct {
			ID   uint
			Time multi_tenants.Time
		}
		err := json.Unmarshal([]byte(timeRep), &v)
		if err != nil {
			t.Errorf("[%d] %s %s", i, timeRep, err.Error())
		}
	}

}
