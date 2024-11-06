package simple

import (
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/Juminiy/kube/pkg/util/zerobuf"
	"github.com/spf13/cast"
	"reflect"
	"time"
)

func Struct(v any) string {
	indirtv := indir(v)
	fieldtyp := indirtv.StructFieldsType()
	fieldtag := indirtv.StructParseTagKV(tagJSON)

	jsonbuf := zerobuf.Get(zerobuf.Small)
	defer jsonbuf.Free()

	jsonbuf.WriteByte('{')
	count := 0
	for name, typ := range fieldtyp {
		var keystr, valstr string
		var valany = indirtv.StructFieldValue(name)
		var strlike = util.AssertStringLike(valany)
		switch {
		case cancompr(typ) || strlike:
			valstr = cast.ToString(valany)

		default:
			logErr(errUnsupportedType)
		}

		keystr = name
		if tkey := fieldtag[name]["key"]; len(tkey) > 0 {
			keystr = tkey
		}

		if len(valstr) == 0 && len(fieldtag[name]["omitempty"]) > 0 {
			continue
		}

		jsonbuf.WriteByte('"')
		jsonbuf.WriteString(keystr)
		jsonbuf.WriteByte('"')
		jsonbuf.WriteByte(':')

		if typ.Kind() == tString || strlike {
			jsonbuf.WriteByte('"')
			jsonbuf.WriteString(valstr)
			jsonbuf.WriteByte('"')
		} else {
			jsonbuf.WriteString(valstr)
		}

		if len(fieldtyp)-1 != count {
			jsonbuf.WriteByte(',')
		}
		count++

	}
	jsonbuf.WriteByte('}')

	return jsonbuf.String()
}

var indir = safe_reflect.IndirectOf

var cancompr = safe_reflect.CanDirectCompare

const (
	tagJSON = "json"
)

const (
	tInvalid reflect.Kind = iota
	tBool
	tInt
	tI8
	tI16
	tI32
	tI64
	tUint
	tU8
	tU16
	tU32
	tU64
	tUPtr
	tF32
	tF64
	tC64
	tC128
	tArr
	tChan
	tFunc
	tAny
	tMap
	tPtr
	tSlice
	tString
	tStruct
	tUnsafePtr
)

var errUnsupportedType = errors.New("unsupported type")
var logErr = stdlog.Error

var timeTyp = indir(time.Now()).Typ
