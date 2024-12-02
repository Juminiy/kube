package safe_json

import (
	"github.com/Juminiy/kube/pkg/util"
	"io"
)

// json safe
// buffer safe
// reflect safe

func SafeMarshal(v any) ([]byte, error) {
	return stdMarshal(v)
}

func SafeMarshalPretty(v any) ([]byte, error) {
	return stdMarshalPretty(v)
}

func SafeUnmarshal(b []byte, v any) error {
	return stdUnmarshal(b, v)
}

func SafeEncoder(wr io.Writer) util.JSONEncoder {
	return stdEncoder(wr)
}

func SafeDecoder(rd io.Reader) util.JSONDecoder {
	return stdDecoder(rd)
}

func From(s string, v any) {
	_ = SafeUnmarshal(util.String2BytesNoCopy(s), v)
}

func String(v any) string {
	bs, _ := SafeMarshal(v)
	return util.Bytes2StringNoCopy(bs)
}

func Bytes(v any) []byte {
	bs, _ := SafeMarshal(v)
	return bs
}

func Pretty(v any) string {
	bs, _ := SafeMarshalPretty(v)
	return util.Bytes2StringNoCopy(bs)
}
