package safe_json

import (
	"bytes"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/elliotchance/orderedmap/v2"
	"github.com/valyala/fastjson"
)

var jsonPool fastjson.ArenaPool

// Expansion
// parse ob to ov
// expand ov to nv
// marshal nv to nb
type Expansion struct {
	ob []byte
	ov *fastjson.Value
	nv *fastjson.Value
	nb []byte
	ar *fastjson.Arena
	vi *orderedmap.OrderedMap[string, *fastjson.Value]
	cg *expansionConfig
	_  *expansionConfig
}

type expansionConfig struct {
	ignoreNull bool
}

type ExpansionOption func(e *Expansion)

func WithIgnoreNull() ExpansionOption {
	return func(e *Expansion) {
		e.cg.ignoreNull = true
	}
}

//type expansionKv struct {
//	//keyBuf   zerobuf.String    // append only key buf
//	//keyFull  string            // a.b.c.d.e copy from keyBuf
//	//keyShort string            // e suffix of keyFull
//	valArray *fastjson.Value // valArray item *fastjson.Value Type is in (0,3,4,5,6,7)
//}

func (e *Expansion) Marshal() []byte {
	return e.nb
}

func ExpandFromBytes(b []byte, o ...ExpansionOption) *Expansion {
	ov, err := fastjson.ParseBytes(b)
	if err != nil {
		return nil
	}
	expansion := &Expansion{
		ob: b,
		ov: ov,
		ar: jsonPool.Get(),
		vi: orderedmap.NewOrderedMap[string, *fastjson.Value](),
		cg: &expansionConfig{},
	}
	for i := range o {
		o[i](expansion)
	}
	expansion.expand()
	jsonPool.Put(expansion.ar)
	return expansion
}

func (e *Expansion) expand() *Expansion {
	e.nv = e.ar.NewArray()
	e.dfs(e.ov, nil)
	for index, el := 0, e.vi.Front(); el != nil; index, el = index+1, el.Next() {
		obj := e.ar.NewObject()
		elKey := keyNoDot0(util.String2BytesNoCopy(el.Key))
		obj.Set("full_key", e.ar.NewStringBytes(elKey))
		obj.Set("short_key", e.ar.NewStringBytes(keyShort(elKey)))
		obj.Set("val_array", el.Value)
		e.nv.SetArrayItem(index, obj)
	}
	e.nb = e.nv.MarshalTo(make([]byte, 0, len(e.ob)))
	return e
}

func (e *Expansion) dfs(src *fastjson.Value, key []byte) {
	switch src.Type() {
	case fastjson.TypeObject:
		srcObj, _ := src.Object()
		srcObj.Visit(func(k []byte, v *fastjson.Value) {
			e.dfs(v, keyDotKey(key, k))
		})

	case fastjson.TypeArray:
		srcArr, _ := src.Array()
		for i := range srcArr {
			e.dfs(srcArr[i], key)
		}

	case fastjson.TypeNull:
		if e.cg.ignoreNull {
			return
		} else {
			e.appendValue(string(key), src)
		}

	//case fastjson.TypeString:
	//case fastjson.TypeNumber:
	//case fastjson.TypeTrue:
	//case fastjson.TypeFalse:
	default:
		e.appendValue(string(key), src)

	}
}

func (e *Expansion) appendValue(key string, value *fastjson.Value) {
	if len(key) == 0 {
		key = onlyArrayNoObjectKey
	}
	if _, ok := e.vi.Get(key); !ok {
		e.vi.Set(key, e.ar.NewArray())
	}
	values := e.vi.GetElement(key).Value
	valuesArray, _ := values.Array()
	values.SetArrayItem(len(valuesArray), value)
}

// multiple dimension(s) array no key
// [e1,e2]
// [[e1],[e2]]
// [[[e1,e2]]]
// [[[...[e1,e2]...]]]
const onlyArrayNoObjectKey = "^no_object_key$"

func keyDotKey(k0, k1 []byte) []byte {
	return append(k0, append([]byte{'.'}, k1...)...)
}

func keyNoDot0(key []byte) []byte {
	if len(key) > 0 && key[0] == '.' {
		return key[1:]
	}
	return key
}

func keyShort(key []byte) []byte {
	lastDotIndex := bytes.LastIndexByte(key, '.')
	if lastDotIndex == -1 {
		return key
	}
	return key[lastDotIndex+1:]
}
