package zero_reflect

import (
	runtimeutil "github.com/Juminiy/kube/pkg/util/runtime"
	"reflect"
	"sync"
)

type Config struct {
	level int
}

func (c Config) Froze() *frozenConfig {
	return &frozenConfig{
		level: c.level,
		cache: new(sync.Map),
	}
}

var Direct = Config{level: _direct}.Froze()
var NoPointer = Config{level: _noPointer}.Froze()
var StructSpec = Config{level: _structSpec}.Froze()
var MustComparable = Config{level: _mustComparable}.Froze()

type frozenConfig struct {
	level int
	cache *sync.Map
}

func (c *frozenConfig) TypeOf(v any) reflect.Type {
	switch c.level {
	case _noPointer:
		return c.indirectTyp(v)

	case _structSpec:
		return c.structIndirectTyp(v)

	case _mustComparable:
		return c.comparableIndirectTyp(v)

	default:
		return c.directTyp(v)
	}
}

// put v direct type to pool
func (c *frozenConfig) directTyp(v any) reflect.Type {
	vTypKey := runtimeutil.EFaceOf(v).Type()
	vTypCache, ok := c.cache.Load(vTypKey)
	if ok {
		typ, ok := vTypCache.(reflect.Type)
		if ok {
			return typ
		}
	}
	vTyp := reflect.TypeOf(v)
	c.cache.Store(vTypKey, vTyp)
	return vTyp
}

// indirectTyp
// 1. put v NoPointer type to pool
// (1). *...*T -> T
func (c *frozenConfig) indirectTyp(v any) reflect.Type {
	return c.directTyp(v)
}

// structIndirectTyp
// 1. (1). as indirectTyp(1)
// 2. put v must struct type to pool
// (2). *...*Struct 			 -> Struct, (as (1))
// (3). *...*[]*...*T 			 -> []*...*T, (T not Struct)
// (4). *...*[num]*...*T 		 -> [num]*...*T, (T not Struct)
// (5). *...*[]*...*Struct 		 -> Struct
// (6). *...*[num]*...*Struct 	 -> Struct
// (7). map[K]E, E ~ (1),(3),(4) -> (1),(3),(4)
// (8). map[K]E, E ~ (2),(5),(6) -> Struct
func (c *frozenConfig) structIndirectTyp(v any) reflect.Type {
	return c.directTyp(v)
}

// comparableIndirectTyp
// 1. (1). as indirectTyp(1)
// 2. (2,5,6,8) as structIndirectTyp(2,5,6,8)
// 3. put v must comparable type to pool
// (9). any -> T, T must be in (reflect.Kind 1~16, Func, String, Struct)
func (c *frozenConfig) comparableIndirectTyp(v any) reflect.Type {
	return c.directTyp(v)
}
