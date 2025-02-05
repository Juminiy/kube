package safe_json

import (
	"github.com/Juminiy/kube/pkg/util"
	goccyjson "github.com/goccy/go-json"
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
	bs, err := STD().MarshalIndent(v0, util.JSONMarshalPrefix, util.JSONMarshalIndent)
	t.Log(util.Bytes2StringNoCopy(bs), err)
}

func TestSafeJSONIteratorMarshal(t *testing.T) {
	bs, err := SafeMarshalPretty(v0)
	t.Log(util.Bytes2StringNoCopy(bs), err)
}

func TestSafeDecoder(t *testing.T) {
	v1 := util.DeepCopyByJSON(STD(), v0)
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

func TestGoccy(t *testing.T) {
	bs, err := goccyjson.Marshal(v0)
	util.Must(err)
	t.Log(util.Bytes2StringNoCopy(bs))
}

var jsonUnmarshalers = map[string]util.JSONUnmarshaler{
	"stdlib":              STD(),
	"json-iterator/std":   JSONIter(),
	"json-iterator/favor": JSONIterFav(),
	"goccy":               GoCCY(),
	"sonic":               Sonic(),
}

var jsonMarshalers = map[string]util.JSONMarshaler{
	"stdlib":              STD(),
	"json-iterator/std":   JSONIter(),
	"json-iterator/favor": JSONIterFav(),
	"goccy":               GoCCY(),
	"sonic":               Sonic(),
}

var jsonLites = map[string]util.JSONLite{
	"stdlib":              STD(),
	"json-iterator/std":   JSONIter(),
	"json-iterator/favor": JSONIterFav(),
	"goccy":               GoCCY(),
	"sonic":               Sonic(),
}

// json BUG
func TestInt64Overflow(t *testing.T) {
	var ofj = []byte("{\"OFII64\":18446744073709551615, \"OFAI64\":18446744073709551616}")

	for name, unl := range jsonUnmarshalers {
		var ofv struct {
			OFII64 uint64
			OFAI64 any
		}
		err := unl.Unmarshal(ofj, &ofv)
		if err != nil {
			t.Logf("%19s: %v", name, err)
		}
		t.Logf("%19s: {%d, %f}", name, ofv.OFII64, ofv.OFAI64)
	}
}

// json BUG
func TestMapAny(t *testing.T) {
	var maj = []byte("{\"name\": \"my-world\", \"id\": 12345}")
	for name, lite := range jsonLites {
		var mapv map[string]any
		util.Must(lite.Unmarshal(maj, &mapv))
		bs, err := lite.Marshal(mapv)
		util.Must(err)
		t.Logf("%19s: %s", name, string(bs))
	}
}
