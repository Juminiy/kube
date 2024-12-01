package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"reflect"
	"strings"
	"testing"
)

type t0 struct {
	I0   int               `valid:"not_zero;range:-10~10;enum:1,2,3"`
	F0   float64           `valid:"not_zero;range:-0.1~0.1;enum:-0.01,0.01,0.09"`
	S0   string            `valid:"not_zero;len:1~10;rule:number"`
	IPtr *int              `valid:"not_nil;range:~2;enum:3,2,1"`
	SPtr *string           `valid:"not_nil;len:10~20;enum:a,c,b"`
	Arr0 []int             `valid:"not_zero;len:~10"`
	Map0 map[string]string `valid:"not_zero;len:~20"`
	I1   int               `valid:"default:1"`
}

func TestDefaultStructE1(t *testing.T) {
	v0 := &t0{
		I0:   1,
		F0:   0.09,
		S0:   "12345",
		IPtr: util.New(3),
		SPtr: util.New("a"),
		Arr0: []int{1, 2, 3},
		Map0: map[string]string{"r": "v"},
		I1:   2,
	}
	t.Log(Default().StructE(v0), v0)
}

func TestDefaultStructE2(t *testing.T) {
	v0 := &t0{
		I0:   111,
		F0:   0,
		S0:   "666",
		IPtr: nil,
		SPtr: nil,
		Arr0: []int{},
		Map0: nil,
	}
	t.Log(Default().StructE(v0), v0)
}

func TestStringSplit(t *testing.T) {
	for _, s := range []string{
		"10",
		"~20",
		"30~",
		"20~60",
		"-1~100",
		"-5~-10",
		"11~2"} {
		t.Log(len(strings.Split(s, "~")), strings.Split(s, "~"))
	}
}

func TestNilNeqNil(t *testing.T) {
	//t.Log(nil == nil)
}

type t1 struct {
	I0 *int    `valid:"enum:1,2,3"`
	I1 *int    `valid:"range:10~100"`
	S0 *string `valid:"len:5~10"`
	S1 *string `valid:"not_nil;len:1~3"`
	S2 *int    `valid:"default:5"`
}

func TestStrictStructE5(t *testing.T) {
	t.Log(Strict().StructE(
		t1{
			I0: util.New(1),
			I1: util.New(77),
			S0: util.New("abcde"),
			S1: util.New("中国"),
			S2: util.New(0),
		}))
}

func TestStrictStructE1(t *testing.T) {
	t.Log(Strict().StructE(
		t1{
			I0: util.New(19),
			I1: util.New(2),
			S0: util.NewString("rrr"),
		},
	))
}

func TestStrictStructE2(t *testing.T) {
	t.Log(Strict().StructE(
		t1{
			I0: nil,
			I1: nil,
			S0: nil,
		},
	))
}

func TestStrictStructE3(t *testing.T) {
	t.Log(Strict().StructE(t1{
		I0: util.New(1),
		I1: util.New(11),
		S0: util.New("xxxxxx"),
		S1: util.New("v"),
	}))
}

func TestStrictStructE4(t *testing.T) {
	t.Log(Strict().StructE(t1{
		I0: util.New(1),
		I1: util.New(11),
		S0: util.New("xxxxxx"),
		S1: util.Zero[*string](),
	}))
}

func TestParseTagK(t *testing.T) {
	tagK := "nil"
	for _, prefix := range []string{
		"",          // 0
		"not_", "!", // 1
		"not_not_", "not_!", "!not_", "!!", // 2
		"not_not_not_", "not_not_!", "not_!not_", "not_!!", "!not_not_", "!not_!", "!!not_", "!!!", // 3
	} {
		t.Logf("%19s -> %7s", prefix+tagK, parseTagK(prefix+tagK))
	}
}

func TestAssignValueToPtr(t *testing.T) {
	var iptr *int
	safe_reflect.Set(666, iptr)
	t.Log(iptr)

	var iptr2 = util.New(777)
	safe_reflect.Set(999, iptr2)
	t.Log(iptr2, *iptr2)

	var iptr3 *int
	var i3 int
	reflect.ValueOf(&i3).Elem().Set(reflect.ValueOf(1024))
	t.Log(i3)
	reflect.ValueOf(&iptr3).Elem().Set(reflect.ValueOf(util.New(888)))
	t.Log(iptr3, *iptr3)
}

func TestJSONAssign(t *testing.T) {
	var v3 struct {
		Name         *string                                                  `json:"name,omitempty"`
		Age          *int                                                     `json:"age,omitempty"`
		Region       **int                                                    `json:"region,omitempty"`
		UnlimitedPtr **************************************************string `json:"unlimited_ptr,omitempty"`
		E0           any                                                      `json:"e0,omitempty"`
		E1           *any                                                     `json:"e1,omitempty"`
	}
	jsonStr := `{"name": "Bob", "age": 18, "region": 6, "unlimited_ptr": "mom", "e0": 999, "e1": "srv6"}`
	safe_json.From(jsonStr, v3)
	t.Log(safe_json.String(v3))

	safe_json.From(jsonStr, &v3)
	t.Log(safe_json.String(v3))
}

func TestStrictStructEDefaultOf(t *testing.T) {
	var v2 struct {
		I0v   int     `valid:"default:333"`
		I0ptr *int    `valid:"default:666"`
		S0    *string `valid:"default:srte"`
		E0    any     `valid:"default:any_string"`
		E1    *any    `valid:"default:any_full"`
	}
	v2.S0 = util.New("srv6")
	t.Log(Strict().StructE(&v2), safe_json.String(v2))
}
