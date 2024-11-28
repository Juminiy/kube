package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util"
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

func TestStruct(t *testing.T) {
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
	t.Log(Default().Struct(v0), v0)
}

func TestStruct2(t *testing.T) {
	v0 := &t0{
		I0:   111,
		F0:   0,
		S0:   "666",
		IPtr: nil,
		SPtr: nil,
		Arr0: []int{},
		Map0: nil,
	}
	t.Log(Default().Struct(v0), v0)
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

func TestConfig_Struct(t *testing.T) {
	t.Log(Strict().StructE(
		struct {
			I0 *int    `valid:"enum:1,2,3"`
			I1 *int    `valid:"range:10~100"`
			S0 *string `valid:"len:5~10"`
		}{
			I0: util.New(19),
			I1: util.New(2),
			S0: util.NewString("rrr"),
		},
	))
}

func TestConfig_Struct2(t *testing.T) {
	t.Log(Strict().StructE(
		struct {
			I0 *int    `valid:"enum:1,2,3"`
			I1 *int    `valid:"range:10~100"`
			S0 *string `valid:"len:5~10"`
		}{
			I0: nil,
			I1: nil,
			S0: nil,
		},
	))
}
