package util

import (
	"encoding/json"
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"reflect"
)

const (
	JSONMarshalPrefix = ""
	JSONMarshalIndent = "  "
)

type StdJSONEncoder interface {
	Marshal(v any) ([]byte, error)
}

type StdJSONDecoder interface {
	Unmarshal(data []byte, v any) error
}

type StdJSON interface {
	StdJSONEncoder
	StdJSONDecoder
}

type JSONEncoder interface {
	Marshal(v any) []byte
}

type JSONDecoder interface {
	Unmarshal(b []byte, v any)
}

type JSONer interface {
	JSONEncoder
	JSONDecoder
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

// DeepCopyByJSON
// no tested yet
func DeepCopyByJSON(stdJSON StdJSON, v any) any {
	bs, encodeErr := stdJSON.Marshal(v)
	if encodeErr != nil {
		stdlog.ErrorF("deepcopy encode value: %v json marshal error: %s", v, encodeErr.Error())
		return nil
	}
	newV := reflect.New(reflect.TypeOf(v)).Interface()
	decodeErr := stdJSON.Unmarshal(bs, &newV)
	if decodeErr != nil {
		stdlog.ErrorF("deepcopy decode value: %v json unmarshal error: %s", newV, decodeErr.Error())
		return nil
	}
	return newV
}
