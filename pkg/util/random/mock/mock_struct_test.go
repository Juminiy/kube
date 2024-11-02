package mock

import (
	"testing"
	"time"
)

type t0 struct {
	ID        uint      `mock:"range:1~1024;"`
	Category  int       `mock:"min:;max:;positive;negative"`
	CreatedAt time.Time `mock:"now;"`
	UpdatedAt time.Time `mock:"range:"`
	DeletedAt time.Time `mock:"min:;max:"`
	Name      string    `mock:"len:1~1024"`
	Desc      string    `mock:"regexp"`
	BusVal0   string    `mock:"uuid;"`
	BusVal1   string    `mock:"alpha"`
	BusVal2   string    `mock:"numeric"`
	BusVal3   string    `mock:"alpha;numeric"`
	BusVal5   string    `mock:"symbol"`
	BusVal6   string    `mock:"enum:str1,str2,str3;"`
	BusVal7   string    `mock:"binary;octal;hexadecimal"`
	BusVal8   string    `mock:"char:2,4,x,q,t,T,d,<,;"`
	BusVal9   string    `mock:"timestamp"`
}

func TestStruct(t *testing.T) {
	v0 := t0{}
	Struct(&v0)
	t.Log(v0)
}
