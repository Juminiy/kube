package safe_json

import (
	stdjson "encoding/json"
	"github.com/Juminiy/kube/pkg/util"
	"io"
)

func stdMarshal(v any) ([]byte, error) {
	return stdjson.Marshal(v)
}

func stdUnmarshal(b []byte, v any) error {
	return stdjson.Unmarshal(b, v)
}

func stdMarshalPretty(v any) ([]byte, error) {
	return stdjson.MarshalIndent(v, util.JSONMarshalPrefix, util.JSONMarshalIndent)
}

func stdEncoder(wr io.Writer) util.JSONEncoder {
	return stdjson.NewEncoder(wr)
}

func stdDecoder(rd io.Reader) util.JSONDecoder {
	return stdjson.NewDecoder(rd)
}

type std struct{}

func (std) Marshal(v any) ([]byte, error) {
	return stdjson.Marshal(v)
}

func (std) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return stdjson.MarshalIndent(v, prefix, indent)
}

func (std) Unmarshal(b []byte, v any) error {
	return stdjson.Unmarshal(b, v)
}

var _std std

func STD() std {
	return _std
}
