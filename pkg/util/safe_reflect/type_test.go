package safe_reflect

import (
	"reflect"
	"testing"
)

func TestStructType(t *testing.T) {
	t.Log(StructType(struct{}{}))
	t.Log(StructType(map[string]struct{}{}))
	t.Log(StructType([]struct{}{}))
	t.Log(StructType([4]struct{}{}))

	t.Log(StructType(&struct{}{}))
	t.Log(StructType(&map[string]*struct{}{}))
	t.Log(StructType(&[]*struct{}{}))
	t.Log(StructType(&[4]*struct{}{}))

	t.Log(StructType(int(3)))
	t.Log(StructType(&map[string]*int{}))
	t.Log(StructType(&[]*int{}))
	t.Log(StructType(&[4]*int{}))
}

func TestStructGetTag2(t *testing.T) {
	type structTyp struct {
		Name string `gorm:"column:name" app:"field:af"`
	}
	t.Log(StructGetTag2(structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2([]structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2([5]structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2(map[string]structTyp{}, "gorm", "column", "app", "field"))

	t.Log(StructGetTag2(&structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2(&[]*structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2(&[5]*structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2(&map[string]*structTyp{}, "gorm", "column", "app", "field"))

	t.Log(StructGetTag2(&[][]*structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2(&[]*[]*[]*structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2(&[5]*[]*[10]*structTyp{}, "gorm", "column", "app", "field"))
	t.Log(StructGetTag2(&map[string][]map[string]map[string]*structTyp{}, "gorm", "column", "app", "field"))
}

func TestTypeFor(t *testing.T) {
	t.Log(reflect.TypeFor[int]().Name())
}

func TestTypeAlias(t *testing.T) {
	type tStruct struct {
		I int
	}
	type t0Struct tStruct
	type t1Struct = tStruct
	t.Log(reflect.TypeFor[tStruct]().String())  // tStruct
	t.Log(reflect.TypeFor[t0Struct]().String()) // t0Struct
	t.Log(reflect.TypeFor[t1Struct]().String()) // tStruct
}
