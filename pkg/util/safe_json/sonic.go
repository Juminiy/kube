package safe_json

import (
	sonicjson "github.com/bytedance/sonic"
)

type sonic struct{}

func (sonic) Default() sonicjson.API {
	return sonicjson.ConfigDefault
}

func (sonic) Std() sonicjson.API {
	return sonicjson.ConfigStd
}

func (sonic) Fast() sonicjson.API {
	return sonicjson.ConfigFastest
}

func (s sonic) Marshal(v any) ([]byte, error) {
	return s.Fast().Marshal(v)
}

func (s sonic) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return s.Fast().MarshalIndent(v, prefix, indent)
}

func (s sonic) Unmarshal(b []byte, v any) error {
	return s.Fast().Unmarshal(b, v)
}

var _Sonic sonic

func Sonic() sonic {
	return _Sonic
}
