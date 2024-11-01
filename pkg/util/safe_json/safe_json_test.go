package safe_json

import (
	"encoding/json"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random/mock"
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
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Desc      string
	Category  int
	BusVal0   string
	BusVal1   string
	BusVal2   string
	Bus1Val0  uint
	Bus2Val1  uint
}

func TestSafeEncoder(t *testing.T) {
	t0Slice := make([]t0, 32)
	mock.Slice(&t0Slice)
	bs, _ := SafeMarshal(t0Slice)
	t.Log(len(bs))
}
