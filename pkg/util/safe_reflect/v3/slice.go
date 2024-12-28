package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
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

func (v V) SliceCallMethod(name string, args []any) (rets []any, called bool) {
	if v.Kind() != reflect.Slice ||
		v.Len() == 0 {
		return nil, false
	}
	return WrapV(v.Index(0)).Indirect().CallMethod(name, args)
}

func (v V) SliceStructSetField(index int, nv map[string]any) {
	if v.Kind() != reflect.Slice || !v.CanSet() || v.Len() <= index {
		return
	}
	WrapV(v.Index(index)).StructSet(nv)
}

func (v V) SliceSet(index int, i any) {
	if v.Kind() != reflect.Slice || !v.CanSet() || v.Len() <= index {
		return
	}
	V2Wrap(v.Index(index)).SetILike(i)
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
	V2Wrap(v.Index(oldLen)).SetILike(i)
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
	if t.Kind() != reflect.Slice {
		return nil
	}
	values := make([]map[string]any, v.Len())
	for i := range v.Len() {
		values[i] = Wrap(v.Index(i)).StructValues()
	}
	return values
}
