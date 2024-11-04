package safe_json

import (
	"encoding/json"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
	"time"
)

var v0 = struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Latitude  float64   `json:"latitude"`
	Canonical bool      `json:"canonical"`
	//Angle     complex128 `json:"angle"`
	IFace   any   `json:"i_face"`
	UintPtr *uint `json:"uint_ptr"`
}{
	ID:        1,
	CreatedAt: time.Now(),
	Name:      "Neko",
	Latitude:  6.66,
	Canonical: true,
	//Angle:     complex(22.33, 114.514),
	IFace:   int8(3),
	UintPtr: util.New[uint](3),
}

func TestStdJSONMarshal(t *testing.T) {
	bs, err := json.MarshalIndent(v0, util.JSONMarshalPrefix, util.JSONMarshalIndent)
	t.Log(util.Bytes2StringNoCopy(bs), err)
}

func TestSafeJSONIteratorMarshal(t *testing.T) {
	bs, err := SafeMarshalPretty(v0)
	t.Log(util.Bytes2StringNoCopy(bs), err)
}

func TestSafeDecoder(t *testing.T) {
	v1 := util.DeepCopyByJSON(safeConfig, v0)
	t.Log(v1)
}

type t0 struct {
	ID        uint      `mock:"range:1~1000"`
	CreatedAt time.Time `mock:"now"`
	UpdatedAt time.Time `mock:"null"`
	DeletedAt time.Time `mock:"null"`
	Name      string    `mock:"len:1~16"`
	Desc      string    `mock:"len:16~32"`
	Category  int       `mock:"enum:1,2,3"`
	BusVal0   string    `mock:"uuid"`
	BusVal1   string    `mock:"len:1~128;alpha;numeric"`
	BusVal2   string    `mock:"len:64~128;symbol"`
	Bus1Val0  uint      `mock:"len:1024~1444;alpha"`
	Bus2Val1  uint      `mock:"len:1000;char:<,>,?"`
}

func TestSafeEncoder(t *testing.T) {
	t0Slice := make([]t0, 32)
	//mock.Slice(&t0Slice)
	//t.Log(Pretty(t0Slice))
	t.Log(len(String(t0Slice))) // 12 field mixed-type, slice len 32: size: 15000B ~ 14KiB
}
