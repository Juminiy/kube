package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestStructEnumOfR(t *testing.T) {
	type t2 struct {
		B0      bool              `valid:"!enum:true,false"`
		I0      int               `valid:"!enum:1,10,25"`
		I1      int8              `valid:"!enum:1,10,25"`
		I2      int16             `valid:"!enum:1,10,25"`
		I3      int32             `valid:"!enum:1,10,25"`
		I4      int64             `valid:"!enum:1,10,25"`
		U0      uint              `valid:"!enum:1,10,25"`
		U1      uint8             `valid:"!enum:1,10,25"`
		U2      uint16            `valid:"!enum:1,10,25"`
		U3      uint32            `valid:"!enum:1,10,25"`
		U4      uint64            `valid:"!enum:1,10,25"`
		U5      uintptr           `valid:"!enum:1,10,25"`
		F0      float32           `valid:"!enum:1,10,25"`
		F1      float64           `valid:"!enum:1,10,25"`
		S0      string            `valid:"!enum:1,10,25"`
		Arr0    [3]int            `valid:"!enum:1,10,25"`
		A0      any               `valid:"!enum:1,10,25"`
		M0      map[string]string `valid:"!enum:1,10,25"`
		Slice0  []int             `valid:"!enum:1,10,25"`
		Struct0 t1                `valid:"!enum:1,10,25"`
	}
	type t2ptr struct {
		B0      *bool              `valid:"!enum:true,false"`
		I0      *int               `valid:"!enum:1,10,25"`
		I1      *int8              `valid:"!enum:1,10,25"`
		I2      *int16             `valid:"!enum:1,10,25"`
		I3      *int32             `valid:"!enum:1,10,25"`
		I4      *int64             `valid:"!enum:1,10,25"`
		U0      *uint              `valid:"!enum:1,10,25"`
		U1      *uint8             `valid:"!enum:1,10,25"`
		U2      *uint16            `valid:"!enum:1,10,25"`
		U3      *uint32            `valid:"!enum:1,10,25"`
		U4      *uint64            `valid:"!enum:1,10,25"`
		U5      *uintptr           `valid:"!enum:1,10,25"`
		F0      *float32           `valid:"!enum:1,10,25"`
		F1      *float64           `valid:"!enum:1,10,25"`
		S0      *string            `valid:"!enum:1,10,25"`
		Arr0    *[3]int            `valid:"!enum:1,10,25"`
		A0      *any               `valid:"!enum:1,10,25"`
		M0      *map[string]string `valid:"!enum:1,10,25"`
		Slice0  *[]int             `valid:"!enum:1,10,25"`
		Struct0 *t1                `valid:"!enum:1,10,25"`
	}

	v2 := t2{
		B0:      true,
		I0:      10,
		I1:      10,
		I2:      10,
		I3:      10,
		I4:      10,
		U0:      10,
		U1:      10,
		U2:      10,
		U3:      10,
		U4:      10,
		U5:      10,
		F0:      10,
		F1:      10,
		S0:      "10",
		Arr0:    [3]int{1, 10, 25},
		A0:      any(10),
		M0:      map[string]string{"1": "", "10": "", "25": ""},
		Slice0:  []int{1, 10, 25},
		Struct0: t1{},
	}
	v2ptr := t2ptr{
		B0:      util.New(true),
		I0:      util.New[int](10),
		I1:      util.New[int8](10),
		I2:      util.New[int16](10),
		I3:      util.New[int32](10),
		I4:      util.New[int64](10),
		U0:      util.New[uint](10),
		U1:      util.New[uint8](10),
		U2:      util.New[uint16](10),
		U3:      util.New[uint32](10),
		U4:      util.New[uint64](10),
		U5:      util.New[uintptr](10),
		F0:      util.New[float32](10),
		F1:      util.New[float64](10),
		S0:      util.New[string]("10"),
		Arr0:    util.New([3]int{1, 10, 25}),
		A0:      util.New(any(10)),
		M0:      util.New(map[string]string{"1": "", "10": "", "25": ""}),
		Slice0:  util.New([]int{1, 10, 25}),
		Struct0: util.New(t1{}),
	}

	t.Log(Strict().StructE(v2))
	t.Log(Strict().StructE(v2ptr))
}

func TestStructLenOfR(t *testing.T) {
	type t2 struct {
		B0      bool              `valid:"!len:1~3"`
		I0      int               `valid:"!len:1~3"`
		I1      int8              `valid:"!len:1~3"`
		I2      int16             `valid:"!len:1~3"`
		I3      int32             `valid:"!len:1~3"`
		I4      int64             `valid:"!len:1~3"`
		U0      uint              `valid:"!len:1~3"`
		U1      uint8             `valid:"!len:1~3"`
		U2      uint16            `valid:"!len:1~3"`
		U3      uint32            `valid:"!len:1~3"`
		U4      uint64            `valid:"!len:1~3"`
		U5      uintptr           `valid:"!len:1~3"`
		F0      float32           `valid:"!len:1~3"`
		F1      float64           `valid:"!len:1~3"`
		S0      string            `valid:"!len:1~3"`
		Arr0    [3]int            `valid:"!len:1~3"`
		A0      any               `valid:"!len:1~3"`
		M0      map[string]string `valid:"!len:1~3"`
		Slice0  []int             `valid:"!len:1~3"`
		Struct0 t1                `valid:"!len:1~3"`
	}
	type t2ptr struct {
		B0      *bool              `valid:"!len:1~3"`
		I0      *int               `valid:"!len:1~3"`
		I1      *int8              `valid:"!len:1~3"`
		I2      *int16             `valid:"!len:1~3"`
		I3      *int32             `valid:"!len:1~3"`
		I4      *int64             `valid:"!len:1~3"`
		U0      *uint              `valid:"!len:1~3"`
		U1      *uint8             `valid:"!len:1~3"`
		U2      *uint16            `valid:"!len:1~3"`
		U3      *uint32            `valid:"!len:1~3"`
		U4      *uint64            `valid:"!len:1~3"`
		U5      *uintptr           `valid:"!len:1~3"`
		F0      *float32           `valid:"!len:1~3"`
		F1      *float64           `valid:"!len:1~3"`
		S0      *string            `valid:"!len:1~3"`
		Arr0    *[3]int            `valid:"!len:1~3"`
		A0      *any               `valid:"!len:1~3"`
		M0      *map[string]string `valid:"!len:1~3"`
		Slice0  *[]int             `valid:"!len:1~3"`
		Struct0 *t1                `valid:"!len:1~3"`
	}

	v2 := t2{
		B0:      true,
		I0:      10,
		I1:      10,
		I2:      10,
		I3:      10,
		I4:      10,
		U0:      10,
		U1:      10,
		U2:      10,
		U3:      10,
		U4:      10,
		U5:      10,
		F0:      10,
		F1:      10,
		S0:      "10",
		Arr0:    [3]int{1, 10, 25},
		A0:      any(10),
		M0:      map[string]string{"1": "", "10": "", "25": ""},
		Slice0:  []int{1, 10, 25},
		Struct0: t1{},
	}
	v2ptr := t2ptr{
		B0:      util.New(true),
		I0:      util.New[int](10),
		I1:      util.New[int8](10),
		I2:      util.New[int16](10),
		I3:      util.New[int32](10),
		I4:      util.New[int64](10),
		U0:      util.New[uint](10),
		U1:      util.New[uint8](10),
		U2:      util.New[uint16](10),
		U3:      util.New[uint32](10),
		U4:      util.New[uint64](10),
		U5:      util.New[uintptr](10),
		F0:      util.New[float32](10),
		F1:      util.New[float64](10),
		S0:      util.New[string]("10"),
		Arr0:    util.New([3]int{1, 10, 25}),
		A0:      util.New(any(10)),
		M0:      util.New(map[string]string{"1": "", "10": "", "25": ""}),
		Slice0:  util.New([]int{1, 10, 25}),
		Struct0: util.New(t1{}),
	}

	t.Log(Strict().StructE(v2))
	t.Log(Strict().StructE(v2ptr))
}

func TestStructNilOfR(t *testing.T) {
	type t2 struct {
		B0      bool              `valid:"!not_nil"`
		I0      int               `valid:"!not_nil"`
		I1      int8              `valid:"!not_nil"`
		I2      int16             `valid:"!not_nil"`
		I3      int32             `valid:"!not_nil"`
		I4      int64             `valid:"!not_nil"`
		U0      uint              `valid:"!not_nil"`
		U1      uint8             `valid:"!not_nil"`
		U2      uint16            `valid:"!not_nil"`
		U3      uint32            `valid:"!not_nil"`
		U4      uint64            `valid:"!not_nil"`
		U5      uintptr           `valid:"!not_nil"`
		F0      float32           `valid:"!not_nil"`
		F1      float64           `valid:"!not_nil"`
		S0      string            `valid:"!not_nil"`
		Arr0    [3]int            `valid:"!not_nil"`
		A0      any               `valid:"!not_nil"`
		M0      map[string]string `valid:"!not_nil"`
		Slice0  []int             `valid:"!not_nil"`
		Struct0 t1                `valid:"!not_nil"`
	}
	type t2ptr struct {
		B0      *bool              `valid:"!not_nil"`
		I0      *int               `valid:"!not_nil"`
		I1      *int8              `valid:"!not_nil"`
		I2      *int16             `valid:"!not_nil"`
		I3      *int32             `valid:"!not_nil"`
		I4      *int64             `valid:"!not_nil"`
		U0      *uint              `valid:"!not_nil"`
		U1      *uint8             `valid:"!not_nil"`
		U2      *uint16            `valid:"!not_nil"`
		U3      *uint32            `valid:"!not_nil"`
		U4      *uint64            `valid:"!not_nil"`
		U5      *uintptr           `valid:"!not_nil"`
		F0      *float32           `valid:"!not_nil"`
		F1      *float64           `valid:"!not_nil"`
		S0      *string            `valid:"!not_nil"`
		Arr0    *[3]int            `valid:"!not_nil"`
		A0      *any               `valid:"!not_nil"`
		M0      *map[string]string `valid:"!not_nil"`
		Slice0  *[]int             `valid:"!not_nil"`
		Struct0 *t1                `valid:"!not_nil"`
	}

	v2 := t2{
		B0:      true,
		I0:      10,
		I1:      10,
		I2:      10,
		I3:      10,
		I4:      10,
		U0:      10,
		U1:      10,
		U2:      10,
		U3:      10,
		U4:      10,
		U5:      10,
		F0:      10,
		F1:      10,
		S0:      "10",
		Arr0:    [3]int{1, 10, 25},
		A0:      any(10),
		M0:      map[string]string{"1": "", "10": "", "25": ""},
		Slice0:  []int{1, 10, 25},
		Struct0: t1{},
	}
	v2ptr := t2ptr{
		B0:      util.New(true),
		I0:      util.New[int](10),
		I1:      util.New[int8](10),
		I2:      util.New[int16](10),
		I3:      util.New[int32](10),
		I4:      util.New[int64](10),
		U0:      util.New[uint](10),
		U1:      util.New[uint8](10),
		U2:      util.New[uint16](10),
		U3:      util.New[uint32](10),
		U4:      util.New[uint64](10),
		U5:      util.New[uintptr](10),
		F0:      util.New[float32](10),
		F1:      util.New[float64](10),
		S0:      util.New[string]("10"),
		Arr0:    util.New([3]int{1, 10, 25}),
		A0:      util.New(any(10)),
		M0:      util.New(map[string]string{"1": "", "10": "", "25": ""}),
		Slice0:  util.New([]int{1, 10, 25}),
		Struct0: util.New(t1{}),
	}

	t.Log(Strict().StructE(v2))
	t.Log(Strict().StructE(v2ptr))
}

func TestStructRangeOfR(t *testing.T) {
	type t2 struct {
		B0      bool              `valid:"!len:1~25"`
		I0      int               `valid:"!len:1~25"`
		I1      int8              `valid:"!len:1~25"`
		I2      int16             `valid:"!len:1~25"`
		I3      int32             `valid:"!len:1~25"`
		I4      int64             `valid:"!len:1~25"`
		U0      uint              `valid:"!len:1~25"`
		U1      uint8             `valid:"!len:1~25"`
		U2      uint16            `valid:"!len:1~25"`
		U3      uint32            `valid:"!len:1~25"`
		U4      uint64            `valid:"!len:1~25"`
		U5      uintptr           `valid:"!len:1~25"`
		F0      float32           `valid:"!len:1~25"`
		F1      float64           `valid:"!len:1~25"`
		S0      string            `valid:"!len:1~25"`
		Arr0    [3]int            `valid:"!len:1~25"`
		A0      any               `valid:"!len:1~25"`
		M0      map[string]string `valid:"!len:1~25"`
		Slice0  []int             `valid:"!len:1~25"`
		Struct0 t1                `valid:"!len:1~25"`
	}
	type t2ptr struct {
		B0      *bool              `valid:"!len:1~25"`
		I0      *int               `valid:"!len:1~25"`
		I1      *int8              `valid:"!len:1~25"`
		I2      *int16             `valid:"!len:1~25"`
		I3      *int32             `valid:"!len:1~25"`
		I4      *int64             `valid:"!len:1~25"`
		U0      *uint              `valid:"!len:1~25"`
		U1      *uint8             `valid:"!len:1~25"`
		U2      *uint16            `valid:"!len:1~25"`
		U3      *uint32            `valid:"!len:1~25"`
		U4      *uint64            `valid:"!len:1~25"`
		U5      *uintptr           `valid:"!len:1~25"`
		F0      *float32           `valid:"!len:1~25"`
		F1      *float64           `valid:"!len:1~25"`
		S0      *string            `valid:"!len:1~25"`
		Arr0    *[3]int            `valid:"!len:1~25"`
		A0      *any               `valid:"!len:1~25"`
		M0      *map[string]string `valid:"!len:1~25"`
		Slice0  *[]int             `valid:"!len:1~25"`
		Struct0 *t1                `valid:"!len:1~25"`
	}

	v2 := t2{
		B0:      true,
		I0:      10,
		I1:      10,
		I2:      10,
		I3:      10,
		I4:      10,
		U0:      10,
		U1:      10,
		U2:      10,
		U3:      10,
		U4:      10,
		U5:      10,
		F0:      10,
		F1:      10,
		S0:      "10",
		Arr0:    [3]int{1, 10, 25},
		A0:      any(10),
		M0:      map[string]string{"1": "", "10": "", "25": ""},
		Slice0:  []int{1, 10, 25},
		Struct0: t1{},
	}
	v2ptr := t2ptr{
		B0:      util.New(true),
		I0:      util.New[int](10),
		I1:      util.New[int8](10),
		I2:      util.New[int16](10),
		I3:      util.New[int32](10),
		I4:      util.New[int64](10),
		U0:      util.New[uint](10),
		U1:      util.New[uint8](10),
		U2:      util.New[uint16](10),
		U3:      util.New[uint32](10),
		U4:      util.New[uint64](10),
		U5:      util.New[uintptr](10),
		F0:      util.New[float32](10),
		F1:      util.New[float64](10),
		S0:      util.New[string]("10"),
		Arr0:    util.New([3]int{1, 10, 25}),
		A0:      util.New(any(10)),
		M0:      util.New(map[string]string{"1": "", "10": "", "25": ""}),
		Slice0:  util.New([]int{1, 10, 25}),
		Struct0: util.New(t1{}),
	}

	t.Log(Strict().StructE(v2))
	t.Log(Strict().StructE(v2ptr))
}

func TestStructZeroOfR(t *testing.T) {
	type t2 struct {
		B0      bool              `valid:"!not_zero"`
		I0      int               `valid:"!not_zero"`
		I1      int8              `valid:"!not_zero"`
		I2      int16             `valid:"!not_zero"`
		I3      int32             `valid:"!not_zero"`
		I4      int64             `valid:"!not_zero"`
		U0      uint              `valid:"!not_zero"`
		U1      uint8             `valid:"!not_zero"`
		U2      uint16            `valid:"!not_zero"`
		U3      uint32            `valid:"!not_zero"`
		U4      uint64            `valid:"!not_zero"`
		U5      uintptr           `valid:"!not_zero"`
		F0      float32           `valid:"!not_zero"`
		F1      float64           `valid:"!not_zero"`
		S0      string            `valid:"!not_zero"`
		Arr0    [3]int            `valid:"!not_zero"`
		A0      any               `valid:"!not_zero"`
		M0      map[string]string `valid:"!not_zero"`
		Slice0  []int             `valid:"!not_zero"`
		Struct0 t1                `valid:"!not_zero"`
	}
	type t2ptr struct {
		B0      *bool              `valid:"!not_zero"`
		I0      *int               `valid:"!not_zero"`
		I1      *int8              `valid:"!not_zero"`
		I2      *int16             `valid:"!not_zero"`
		I3      *int32             `valid:"!not_zero"`
		I4      *int64             `valid:"!not_zero"`
		U0      *uint              `valid:"!not_zero"`
		U1      *uint8             `valid:"!not_zero"`
		U2      *uint16            `valid:"!not_zero"`
		U3      *uint32            `valid:"!not_zero"`
		U4      *uint64            `valid:"!not_zero"`
		U5      *uintptr           `valid:"!not_zero"`
		F0      *float32           `valid:"!not_zero"`
		F1      *float64           `valid:"!not_zero"`
		S0      *string            `valid:"!not_zero"`
		Arr0    *[3]int            `valid:"!not_zero"`
		A0      *any               `valid:"!not_zero"`
		M0      *map[string]string `valid:"!not_zero"`
		Slice0  *[]int             `valid:"!not_zero"`
		Struct0 *t1                `valid:"!not_zero"`
	}

	v2 := t2{
		B0:      true,
		I0:      10,
		I1:      10,
		I2:      10,
		I3:      10,
		I4:      10,
		U0:      10,
		U1:      10,
		U2:      10,
		U3:      10,
		U4:      10,
		U5:      10,
		F0:      10,
		F1:      10,
		S0:      "10",
		Arr0:    [3]int{1, 10, 25},
		A0:      any(10),
		M0:      map[string]string{"1": "", "10": "", "25": ""},
		Slice0:  []int{1, 10, 25},
		Struct0: t1{},
	}
	v2ptr := t2ptr{
		B0:      util.New(true),
		I0:      util.New[int](10),
		I1:      util.New[int8](10),
		I2:      util.New[int16](10),
		I3:      util.New[int32](10),
		I4:      util.New[int64](10),
		U0:      util.New[uint](10),
		U1:      util.New[uint8](10),
		U2:      util.New[uint16](10),
		U3:      util.New[uint32](10),
		U4:      util.New[uint64](10),
		U5:      util.New[uintptr](10),
		F0:      util.New[float32](10),
		F1:      util.New[float64](10),
		S0:      util.New[string]("10"),
		Arr0:    util.New([3]int{1, 10, 25}),
		A0:      util.New(any(10)),
		M0:      util.New(map[string]string{"1": "", "10": "", "25": ""}),
		Slice0:  util.New([]int{1, 10, 25}),
		Struct0: util.New(t1{}),
	}

	t.Log(Strict().StructE(v2))
	t.Log(Strict().StructE(v2ptr))
}

func TestStructZeroOfRBug(t *testing.T) {
	type t2 struct {
		I0 int `valid:"zero"`
		I1 int `valid:"not_zero"`
	}
	v2 := t2{
		I0: -1,
		I1: 0,
	}
	t.Log(Strict().StructE(v2))
}
