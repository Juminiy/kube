package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"iter"
	"reflect"
)

func (t T) SliceStructTags(tagKey string) Tags {
	if t.Kind() != reflect.Slice {
		return nil
	}
	return t.Indirect().StructTags(tagKey)
}

func (t T) SliceStructTypes() Types {
	if t.Kind() != reflect.Slice {
		return nil
	}
	return t.Indirect().StructTypes()
}

func (t T) SliceStructFields() Fields {
	if t.Kind() != reflect.Slice {
		return nil
	}
	return t.Indirect().StructFields()
}

func (t T) SliceElemNew() Tv {
	if t.Kind() != reflect.Slice {
		return Tv{}
	}
	return WrapT(t.Elem()).New()
}

func (v V) SliceCallMethod(name string, args []any) (rets []any, called bool) {
	if v.Kind() != reflect.Slice ||
		v.Len() == 0 {
		return nil, false
	}
	return WrapV(v.Index(0)).CallMethod(name, args)
}

func (v V) SliceStructSetField(index int, nv map[string]any) {
	if v.Kind() != reflect.Slice || v.Len() <= index {
		return
	}
	WrapVI(v.Index(index)).StructSet(nv)
}

func (v V) SliceStructSetField2(nv map[string]any) {
	if v.Kind() != reflect.Slice {
		return
	}
	for i := range v.Len() {
		WrapVI(v.Index(i)).StructSet(nv)
	}
}

func (v V) SliceSet(index int, i any) {
	if v.Kind() != reflect.Slice || v.Len() <= index {
		return
	}
	WrapV(v.Index(index)).SetI(i)
}

func (v V) SliceAppend(i any) {
	if v.Kind() != reflect.Slice || !v.CanSet() {
		return
	}
	oldLen, oldCap := v.Len(), v.Cap()
	newLen, newCap := sliceLenCap(oldLen, oldCap)
	v.Grow(newCap - oldCap)
	/*if oldLen+newCap-oldCap > oldCap {
		fmt.Printf("grow happend: %d->%d\n", oldCap, v.Cap())
	}*/
	v.SetLen(newLen)
	WrapV(v.Index(oldLen)).SetI(i)
}

func sliceLenCap(oldLen, oldCap int) (newLen, newCap int) {
	if oldLen <= 0 {
		oldLen = 0
	}
	if oldCap <= 0 {
		newCap = util.MagicSliceCap
	} else {
		if oldCap == oldLen {
			newCap = oldCap * 2
		} else {
			newCap = oldCap
		}
	}
	newLen = oldLen + 1
	if newCap < newLen {
		newCap = newLen
	}
	return
}

func (tv Tv) SliceStructValues() []map[string]any {
	t, v := tv.T, tv.V
	if t.Kind() != reflect.Slice || t.Indirect().Kind() != reflect.Struct {
		return nil
	}
	values := make([]map[string]any, v.Len())
	tv.IterIndex()(func(i int, rv reflect.Value) bool {
		values[i] = WrapI(rv).StructValues()
		return true
	})
	return values
}

func (tv Tv) SliceMapValues() []map[string]any {
	t, v := tv.T, tv.V
	if t.Kind() != reflect.Slice || t.Indirect().Kind() != reflect.Map {
		return nil
	}
	values := make([]map[string]any, v.Len())
	tv.IterIndex()(func(i int, rv reflect.Value) bool {
		values[i] = WrapI(rv).MapValues()
		return true
	})
	return values
}

func (v V) IterIndex() iter.Seq2[int, reflect.Value] {
	if !util.ElemIn(v.Kind(),
		reflect.Slice, reflect.Array) {
		return func(yield func(int, reflect.Value) bool) {
			return
		}
	}
	return func(yield func(int, reflect.Value) bool) {
		for i := range v.Len() {
			if !yield(i, v.Index(i)) {
				return
			}
		}
	}
}
