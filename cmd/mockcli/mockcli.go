package main

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util/random/mock"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"time"
)

func main() {
	mock.Default = true
	t0sl := make([]T0, 32)
	mock.Slice(t0sl)
	stdlog.Info(len(safe_json.String(t0sl)))
}

type T0 struct {
	ID        uint      `mock:"range:1~1024;"`
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