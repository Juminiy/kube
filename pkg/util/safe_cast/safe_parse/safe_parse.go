package safe_parse

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
	"time"
)

type Type interface {
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
	case reflect.String:
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
	readV.numberV = util.New(ParseNumber(s))

	// time.Time, Type.Time()
	timeV, ok := ParseTime(s)
	if ok {
		readV.timeV = util.New(timeV)
	}

	return readV
}
