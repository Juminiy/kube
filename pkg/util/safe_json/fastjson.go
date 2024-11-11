package safe_json

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/valyala/fastjson"
)

type Flatten struct {
	v *fastjson.Value
}

func FlattenFromBytes(b []byte) *Flatten {
	embedv, err := fastjson.ParseBytes(b)
	if err != nil {
		return nil
	}
	return &Flatten{v: flattenFromValue(embedv)}
}

func flattenFromValue(v *fastjson.Value) *fastjson.Value {
	var dst = &fastjson.Value{}
	flattenv(v, dst)
	return dst
}

func flattenv(src, dst *fastjson.Value) {
	switch src.Type() {
	case fastjson.TypeObject:
		//srcObj, _ := src.Object()
		//srcObj.Visit(func(key []byte, v *fastjson.Value) {
		//	dst.Set()
		//})

	case fastjson.TypeArray:

	case fastjson.TypeString:

	case fastjson.TypeNumber:

	case fastjson.TypeTrue:

	case fastjson.TypeFalse:

	case fastjson.TypeNull:

	}
}

func (f Flatten) Marshal() []byte {
	return f.v.MarshalTo(make([]byte, 0, util.Ki))
}
