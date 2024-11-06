// Package mockv2 was generated
package mockv2

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
	"time"
)

type t0 struct {
	ID        uint      `mock:"range:1~10000000;"`
	CreatedAt time.Time `mock:"now;"`
	UpdatedAt time.Time `mock:"null;"`
	DeletedAt time.Time `mock:"null;"`
	Category  int       `mock:"enum:1,2,3"`
	Name      string    `mock:"len:1~32"`
	Desc      string    `mock:"regexp:'[012]*'"`
	BusVal0   string    `mock:"uuid;"`
	BusVal1   string    `mock:"alpha"`
	BusVal2   string    `mock:"numeric"`
	BusVal3   string    `mock:"alpha;numeric"`
	BusVal5   string    `mock:"symbol"`
	BusVal6   string    `mock:"enum:str1,str2,str3;"`
	BusVal7   string    `mock:"binary;octal;hexadecimal"`
	BusVal8   string    `mock:"char:2,4,x,q,t,T,d,<,;"`
	BusVal9   string    `mock:"timestamp"`
	Latitude  float32   `mock:"range:1~1024"`
	Longitude float64   `mock:"range:-9~22"`
}

func TestStruct(t *testing.T) {
	util.Recover(func() {
		for range 1000 {
			v0 := t0{}
			Struct(&v0)
			//t.Log(i)
		}
	})
}
