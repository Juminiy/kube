package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"reflect"
)

type T struct {
	reflect.Type
}

func NewT(i any) T {
	return WrapT(reflect.TypeOf(i))
}

func WrapT(rt reflect.Type) T {
	return T{Type: rt}
}

func (t T) CanElem() bool {
	return util.ElemIn(t.Kind(),
		reflect.Array, reflect.Chan, reflect.Map, reflect.Pointer, reflect.Slice)
}

func (t T) Indirect() T {
	tCopy := t
	for util.ElemIn(tCopy.Kind(),
		reflect.Array, reflect.Slice, reflect.Chan, reflect.Pointer,
	) {
		tCopy = WrapT(tCopy.Elem())
	}
	return tCopy
}

func (t T) IndirectElem() T {
	tCopy := t
	for util.ElemIn(tCopy.Kind(),
		reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.Pointer,
	) {
		tCopy = WrapT(tCopy.Elem())
	}
	return tCopy
}

func rts(i []any) []reflect.Type {
	return lo.Map(i, func(item any, _ int) reflect.Type {
		return reflect.TypeOf(item)
	})
}

func (t T) Tag1(tagKey string) (raw map[string]string) {
	switch t.Kind() {
	case reflect.Struct:
		return t.StructRawTag(tagKey)

	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Pointer:
		return t.Indirect().StructRawTag(tagKey)

	default: // ignore
		return
	}
}

func (t T) Tag2(tagKey, valKey string) map[string]string {
	return lo.MapValues(t.Tags(tagKey), func(tag Tag, name string) string {
		return util.MapElem(tag, valKey)
	})
}

func (t T) Tag2VName(tagKey, valKey string) map[string]string {
	return util.MapVK(t.Tag2(tagKey, valKey))
}

func (t T) Tags(tagKey string) (tags Tags) {
	switch t.Kind() {
	case reflect.Struct:
		tags = t.StructTags(tagKey)

	case reflect.Array:
		tags = t.ArrayStructTags(tagKey)

	case reflect.Slice:
		tags = t.SliceStructTags(tagKey)

	default: // ignore
	}
	return
}

func (t T) CanSet(st reflect.Type) bool {
	return t.Type.Kind() == reflect.Interface || t.Type == st
}

func (t T) FieldType() Types {
	switch t.Kind() {
	case reflect.Struct:
		return t.StructTypes()

	case reflect.Array:
		return t.ArrayStructTypes()

	case reflect.Slice:
		return t.SliceStructTypes()

	case reflect.Map:
		return Types{MapElemType: t.MapElemType()}

	default:
		return nil
	}
}

func (t T) New() Tv {
	if t.Type == nil {
		return Tv{}
	}
	return Wrap(reflect.New(t.Type))
}

func (t T) NewElem() Tv {
	switch t.Kind() {
	case reflect.Array:
		return t.ArrayElemNew()

	case reflect.Slice:
		return t.SliceElemNew()

	case reflect.Map:
		return t.MapElemNew()

	default:
		return t.New()
	}
}

func (t T) IsEFace() bool {
	return t.Type.Kind() == reflect.Interface && t.Type.NumMethod() == 0
}
