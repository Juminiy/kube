package util

import "encoding/json"

const (
	JSONMarshalPrefix = ""
	JSONMarshalIndent = "  "
)

type StdJSONEncoder interface {
	String() (string, error)
	Bytes() ([]byte, error)
}

type JSONEncoder interface {
	String() string
	Bytes() []byte
	Size() int64
}

func MarshalJSONPretty(v any) (string, error) {
	jsonBytes, err := marshalJSONPretty(&v)
	return Bytes2StringNoCopy(jsonBytes), err
}

func marshalJSONPretty(v any) ([]byte, error) {
	return json.MarshalIndent(&v, JSONMarshalPrefix, JSONMarshalIndent)
}
