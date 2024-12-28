package util

import (
	"encoding/json"
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util/zero_reflect"
	"io"
	"reflect"
)

const (
	JSONMarshalPrefix = ""
	JSONMarshalIndent = "  "
)

type JSONMarshaler interface {
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix, indent string) ([]byte, error)
}

type JSONUnmarshaler interface {
	Unmarshal(data []byte, v any) error
}

type JSONEncoder interface {
	Encode(v any) error
	SetIndent(prefix, indent string)
	SetEscapeHTML(on bool)
}

type JSONDecoder interface {
	Decode(v any) error
	More() bool
	Buffered() io.Reader
	UseNumber()
	DisallowUnknownFields()
}

type JSONer interface {
	JSONMarshaler
	JSONUnmarshaler
	JSONEncoder
	JSONDecoder
}

type JSONLite interface {
	JSONMarshaler
	JSONUnmarshaler
}

type Stringer fmt.Stringer

type Byteser interface {
	Bytes() []byte
}

type Sizer interface {
	Size() int64
}

func MarshalJSONPretty(v any) (string, error) {
	jsonBytes, err := marshalJSONPretty(&v)
	return Bytes2StringNoCopy(jsonBytes), err
}

func marshalJSONPretty(v any) ([]byte, error) {
	return json.MarshalIndent(&v, JSONMarshalPrefix, JSONMarshalIndent)
}

// DeepCopyByJSON
// tested is ok
func DeepCopyByJSON(stdJSON JSONLite, v any) any {
	bs, encodeErr := stdJSON.Marshal(v)
	if encodeErr != nil {
		stdlog.ErrorF("deepcopy encode value: %v json marshal error: %s", v, encodeErr.Error())
		return nil
	}
	newV := reflect.New(zero_reflect.TypeOf(v)).Elem().Interface()
	decodeErr := stdJSON.Unmarshal(bs, &newV)
	if decodeErr != nil {
		stdlog.ErrorF("deepcopy decode value: %v json unmarshal error: %s", newV, decodeErr.Error())
		return nil
	}
	return newV
}
