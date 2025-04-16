package zero_reflect

import (
	"github.com/modern-go/reflect2"
	"reflect"
	"testing"
	"unsafe"
)

func TestReflect2Func(t *testing.T) {
	reflect2.IsNil(nil)
	reflect2.TypeOf(nil)
	reflect2.IsNullable(0)
	reflect2.NoEscape(unsafe.Pointer(new(int)))
	reflect2.PtrOf(114514)
	reflect2.PtrTo(&reflect2.UnsafeSliceType{})
	reflect2.RTypeOf(114514)
	reflect2.Type2(reflect.TypeOf(114514))
	reflect2.TypeOfPtr(10)
	reflect2.UnsafeCastString("Hajimi")
	reflect2.DefaultTypeOfKind(0)
	reflect2.IFaceToEFace(unsafe.Pointer(new(int)))
	reflect2.TypeByName("int")
	reflect2.TypeByPackageName("util", "int")
}

func TestReflect2Config(t *testing.T) {
	reflect2.ConfigUnsafe.TypeOf(10)
	reflect2.ConfigSafe.TypeOf(10)
}

func TestReflect2API(t *testing.T) {

}
