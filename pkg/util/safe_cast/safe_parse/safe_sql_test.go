package safe_parse

import (
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
		newTestRep("2025-01-05 18:33:42", KStdTime),
		newTestRep("rrr12211", KBytes),
		newTestRep("qqwsdasd", KJSONRaw),
		newTestRep("qweqssad", KRawBytes),
		newTestRep("rr12112", KNullString),
		newTestRep("11111", KNullInt64),
		newTestRep("222222", KNullInt32),
		newTestRep("121", KNullInt16),
		newTestRep("88", KNullByte),
		newTestRep("114.51212", KNullFloat64),
		newTestRep("false", KNullBool),
		newTestRep("2025-01-05 18:33:42", KNullTime),
		newTestRep(uuid.NewString(), KUUID),
	} {
		pv, ok := Parse(s.RawStr).Get(s.Kind)
		if ok {
			t.Logf("%19s to %10s(%v)", s.RawStr, reflect.TypeOf(pv), pv)
		} else {
			t.Fatalf("%19s to kind: %10s", s.RawStr, s.Kind.String())
		}
	}
}
