package util

import (
	"encoding/json"
	"fmt"
)

const (
	JSONMarshalPrefix = ""
	JSONMarshalIndent = "  "
)

type StdJSONEncoder interface {
	String() (string, error)
	Bytes() ([]byte, error)
}

type JSONEncoder interface {
	Stringer
	Byteser
	Sizer
}

func MarshalJSONPretty(v any) (string, error) {
	jsonBytes, err := marshalJSONPretty(&v)
	return Bytes2StringNoCopy(jsonBytes), err
}

func marshalJSONPretty(v any) ([]byte, error) {
	return json.MarshalIndent(&v, JSONMarshalPrefix, JSONMarshalIndent)
}

type Stringer fmt.Stringer

type Byteser interface {
	Bytes() []byte
}

type Sizer interface {
	Size() int64
}
