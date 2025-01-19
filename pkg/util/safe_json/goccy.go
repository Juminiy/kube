package safe_json

import goccyjson "github.com/goccy/go-json"

type goccy struct{}

func (goccy) Marshal(v any) ([]byte, error) {
	return goccyjson.Marshal(v)
}

func (goccy) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return goccyjson.MarshalIndent(v, prefix, indent)
}

func (goccy) Unmarshal(b []byte, v any) error {
	return goccyjson.Unmarshal(b, v)
}

var _GoCCY = goccy{}

func GoCCY() goccy {
	return _GoCCY
}
