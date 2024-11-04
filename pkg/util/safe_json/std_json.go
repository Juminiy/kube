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
