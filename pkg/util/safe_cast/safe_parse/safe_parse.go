package safe_parse

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
	"time"
)

type Category int8

const (
	Null = Category(0)
	Bool = Category(1)
	Num  = Category(2)
	Time = Category(3)
	Text = Category(4)
)

type Type interface {
	Category() Category
	Bool() bool
	Number() Number
	Time() time.Time
	Text() string
	Get(kind reflect.Kind, kindDesc ...string) (any, bool)
}

type readable struct {
	boolV   *bool
	numberV *Number
	timeV   *time.Time
	stringV string
}

func (r readable) Category() Category {
	if r.timeV != nil {
		return Time
	} else if r.numberV != nil {
		return Num
	} else if r.boolV != nil {
		return Bool
	}
	return Text
}

func (r readable) Bool() bool {
	if r.boolV != nil {
		return *r.boolV
	}
	return false
}

func (r readable) Number() Number {
	if r.numberV != nil {
		return *r.numberV
	}
	return Number{}
}

func (r readable) Time() time.Time {
	if r.timeV != nil {
		return *r.timeV
	}
	return time.Time{}
}

func (r readable) Text() string {
	return r.stringV
}

func (r readable) Get(kind reflect.Kind, kindDesc ...string) (v any, ok bool) {
	switch kind {
	case reflect.Bool:
		if r.boolV != nil {
			return *r.boolV, true
		}
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		if r.numberV != nil {
			return (*r.numberV).Get(kind)
		}
	case reflect.Interface, reflect.String:
		return r.stringV, true
	default:
		if len(kindDesc) > 0 {
			switch kindDesc[0] {
			case "time", "Time", "TIME", "time.Time",
				"nullTime", "NullTime", "NULLTIME", "sql.NullTime":
				if r.timeV != nil {
					return *r.timeV, true
				}
			}
		}
	}
	return nil, false
}

func Parse(s string) Type {
	readV := readable{}

	// string, reflect.String, Type.Text()
	readV.stringV = s

	// bool, reflect.Bool, Type.Bool()
	boolV, ok := ParseBool(s)
	if ok {
		readV.boolV = util.New(boolV)
	}

	// Number, reflect.Int~reflect.F64, Type.Number()
	numberV := ParseNumber(s)
	for _kind := reflect.Int; _kind <= reflect.Float64; _kind++ {
		if _, ok := numberV.Get(_kind); ok {
			readV.numberV = util.New(numberV)
			break
		}
	}

	// time.Time, Type.Time()
	timeV, ok := ParseTime(s)
	if ok {
		readV.timeV = util.New(timeV)
	}

	return readV
}
