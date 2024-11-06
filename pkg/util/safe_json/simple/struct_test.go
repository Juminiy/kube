package simple

import (
	"github.com/Juminiy/kube/pkg/util/random/mock"
	"testing"
	"time"
)

type t0 struct {
	ID        uint      `mock:"range:1~10000000;" json:"id"`
	CreatedAt time.Time `mock:"now;" json:"created_at"`
	UpdatedAt time.Time `mock:"null;" json:"updated_at"`
	DeletedAt time.Time `mock:"null;" json:"deleted_at"`
	Category  int       `mock:"enum:1,2,3" json:"category"`
	Name      string    `mock:"len:1~32" json:"name"`
	Desc      string    `mock:"regexp:'[012]*'" json:"desc"`
	BusVal0   string    `mock:"uuid;" json:"bus_val_0"`
	BusVal1   string    `mock:"alpha" json:"bus_val_1"`
	BusVal2   string    `mock:"numeric" json:"bus_val_2"`
	BusVal3   string    `mock:"alpha;numeric" json:"bus_val_3"`
	BusVal5   string    `mock:"symbol" json:"bus_val_5"`
	BusVal6   string    `mock:"enum:str1,str2,str3;" json:"bus_val_6"`
	BusVal7   string    `mock:"binary;octal;hexadecimal" json:"bus_val_7"`
	BusVal8   string    `mock:"char:2,4,x,q,t,T,d,<,;" json:"bus_val_8"`
	BusVal9   string    `mock:"timestamp" json:"bus_val_9"`
	Latitude  float32   `mock:"range:1~1024" json:"latitude"`
	Longitude float64   `mock:"range:-9~22" json:"longitude"`
}

func TestStruct(t *testing.T) {
	v0 := t0{}
	mock.Struct(&v0)
	t.Log(Struct(&v0))
}
