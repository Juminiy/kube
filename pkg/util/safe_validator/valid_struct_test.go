package safe_validator

import (
	"strings"
	"testing"
)

type t0 struct {
	I0   int               `valid:"not_zero;range:-10~10;enum:1,2,3"`
	F0   float64           `valid:"not_zero;range:-0.1~0.1;enum:-0.01,0.01,0.09"`
	S0   string            `valid:"not_zero;len:1~10;rule:number"`
	IPtr *int              `valid:"not_nil;range:~2;enum:3,2,1"`
	SPtr *string           `valid:"not_nil;len:10~20;enum:a,c,b"`
	Arr0 []int             `valid:"not_zero;len:10"`
	Map0 map[string]string `valid:"not_zero;len:20"`
}

func init() {
	_debug = true
}

func TestStruct(t *testing.T) {
	t.Log(Struct(t0{
		I0:   1,
		F0:   0.09,
		S0:   "12345",
		IPtr: nil,
		SPtr: nil,
		Arr0: []int{1, 2, 3},
		Map0: nil,
	}))

	t.Log(Struct(t0{
		I0:   111,
		F0:   -9,
		S0:   "12345678901",
		IPtr: nil,
		SPtr: nil,
		Arr0: []int{},
		Map0: nil,
	}))
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
