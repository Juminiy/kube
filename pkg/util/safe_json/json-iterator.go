package safe_json

import (
	"github.com/Juminiy/kube/pkg/util"
	jsoniter "github.com/json-iterator/go"
	"io"
)

// json-iterator is unsafe
var (
	stdConfig  = jsoniter.ConfigCompatibleWithStandardLibrary
	fastConfig = jsoniter.ConfigFastest
	safeConfig = jsoniter.Config{
		IndentionStep:                 0,      // for json pretty
		MarshalFloatWith6Digits:       true,   // low accuracy float
		EscapeHTML:                    false,  // no escape for HTML, because no-need html
		SortMapKeys:                   false,  // no need sorted map keys
		UseNumber:                     true,   // must use number
		DisallowUnknownFields:         false,  // unknown field unmarshal, but not error
		TagKey:                        "json", // default tag is json
		OnlyTaggedField:               false,  // allow exported but no tagged field
		ValidateJsonRawMessage:        false,  // no valid, none-sense valid before unmarshal, return error anyway
		ObjectFieldMustBeSimpleString: true,   // key must string
		CaseSensitive:                 true,   // key must be sensitive, or not will be ambiguity
	}.Froze()
)

func unsafeMarshal(v any) ([]byte, error) {
	return safeConfig.Marshal(v)
}

func unsafeMarshalIndent(v any) ([]byte, error) {
	return safeConfig.MarshalIndent(v, util.JSONMarshalPrefix, util.JSONMarshalIndent)
}

func unsafeUnmarshal(b []byte, v any) error {
	return safeConfig.Unmarshal(b, v)
}

func unsafeEncoder(wr io.Writer) util.JSONEncoder {
	return safeConfig.NewEncoder(wr)
}

func unsafeDecoder(rd io.Reader) util.JSONDecoder {
	return safeConfig.NewDecoder(rd)
}

type jsonIter struct{}

func (jsonIter) Marshal(v any) ([]byte, error) {
	return stdConfig.Marshal(v)
}

func (jsonIter) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return stdConfig.MarshalIndent(v, prefix, indent)
}

func (jsonIter) Unmarshal(b []byte, v any) error {
	return stdConfig.Unmarshal(b, v)
}

var _jsonIter jsonIter

func Jsoniter() jsonIter {
	return _jsonIter
}

func SafeConfig() jsoniter.API {
	return safeConfig
}
