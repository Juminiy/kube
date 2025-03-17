package safe_parse

import (
	"database/sql"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestSafeSql(t *testing.T) {
	type testRep struct {
		RawStr string
		Kind   reflect.Kind
	}
	var newTestRep = func(s string, k reflect.Kind) testRep {
		return testRep{RawStr: s, Kind: k}
	}

	for _, s := range []testRep{
		newTestRep("true", reflect.Bool),
		newTestRep("114514", reflect.Int),
		newTestRep("123", reflect.Int8),
		newTestRep("32135", reflect.Int16),
		newTestRep("33194522", reflect.Int32),
		newTestRep("33789288111", reflect.Int64),
		newTestRep("208911112", reflect.Uint),
		newTestRep("243", reflect.Uint8),
		newTestRep("63233", reflect.Uint16),
		newTestRep("2113414", reflect.Uint32),
		newTestRep("757622352", reflect.Uint64),
		newTestRep("1424166464", reflect.Uintptr),
		newTestRep("44.33", reflect.Float32),
		newTestRep("114514.191810", reflect.Float64),
		newTestRep("rr", reflect.Interface),
		newTestRep("r8", reflect.String),
		newTestRep("2025-01-05 18:33:42", KStdTime), newTestRep("2025-01-05 18:33:42", KStdTimePtr),
		newTestRep("rrr12211", KBytes), newTestRep("rrr12211", KBytesPtr),
		newTestRep("qqwsdasd", KJSONRaw), newTestRep("qqwsdasd", KJSONRawPtr),
		newTestRep(uuid.NewString(), KUUID), newTestRep(uuid.NewString(), KUUIDPtr),
		newTestRep("qweqssad", KRawBytes), newTestRep("qweqssad", KRawBytesPtr),
		newTestRep("rr12112", KNullString), newTestRep("rr12112", KNullStringPtr), newTestRep("rr12112", KUnderlyingNullString),
		newTestRep("11111", KNullInt64), newTestRep("11111", KNullInt64Ptr), newTestRep("11111", KUnderlyingNullInt64),
		newTestRep("222222", KNullInt32), newTestRep("222222", KNullInt32Ptr), newTestRep("222222", KUnderlyingNullInt32),
		newTestRep("121", KNullInt16), newTestRep("121", KNullInt16Ptr), newTestRep("121", KUnderlyingNullInt16),
		newTestRep("88", KNullByte), newTestRep("88", KNullBytePtr), newTestRep("88", KUnderlyingNullByte),
		newTestRep("114.51212", KNullFloat64), newTestRep("114.51212", KNullFloat64Ptr), newTestRep("114.51212", KUnderlyingNullFloat64),
		newTestRep("false", KNullBool), newTestRep("false", KNullBoolPtr), newTestRep("false", KUnderlyingNullBool),
		newTestRep("2025-01-05 18:33:42", KNullTime), newTestRep("2025-01-05 18:33:42", KNullTimePtr), newTestRep("2025-01-05 18:33:42", KUnderlyingNullTime),
	} {
		pv, ok := Parse(s.RawStr).Get(s.Kind)
		if ok {
			t.Logf("%19s to %10s(%v)", s.RawStr, reflect.TypeOf(pv), pv)
		} else {
			t.Fatalf("%19s to kind: %10s", s.RawStr, s.Kind.String())
		}
	}
}

func TestSafeSqlConvertible(t *testing.T) {
	type sid sql.NullString
	type ssid sid
	type s2id sql.NullString
	type ss2id s2id

	t.Log(reflect.TypeOf(sid{}).ConvertibleTo(_NullStringType))
	t.Log(reflect.TypeOf(ssid{}).ConvertibleTo(_NullStringType))
	t.Log(reflect.TypeOf(s2id{}).ConvertibleTo(_NullStringType))
	t.Log(reflect.TypeOf(ss2id{}).ConvertibleTo(_NullStringType))
	t.Log(reflect.TypeOf(sid{}).ConvertibleTo(reflect.TypeOf(sid{})))
	t.Log(reflect.TypeOf(sid{}).ConvertibleTo(reflect.TypeOf(ssid{})))
	t.Log(reflect.TypeOf(sid{}).ConvertibleTo(reflect.TypeOf(s2id{})))
	t.Log(reflect.TypeOf(sid{}).ConvertibleTo(reflect.TypeOf(ss2id{})))
}

func TestSafeSqlUnderlyingRType(t *testing.T) {
	type sid sql.NullString
	if parsedValueRt, ok := Parse("iamhajimi").GetByRT(reflect.TypeOf(sid{})); ok {
		t.Log(parsedValueRt.(sid).String, parsedValueRt.(sid).Valid)
	}
}
