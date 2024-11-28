package util

import (
	"github.com/samber/lo"
	"maps"
)

const (
	MagicMapCap = (iota + 1) << 4
)

// MapMerge
// merge field in src `M`as elem copy to field in dst `M`as elem
func MapMerge[Map ~map[F]M, F comparable, M map[K]V, K comparable, V any](dst, src Map) {
	for field := range dst {
		maps.Copy(dst[field], src[field])
	}
}

// MapEvict
// evict different elem_type map with same key_type map with same key
func MapEvict[Map1 ~map[K]V1, Map2 ~map[K]V2, K comparable, V1 any, V2 any](dst Map1, src Map2) {
	evictKeys := make([]K, 0, len(dst))
	for key := range dst {
		if !MapOk(src, key) {
			evictKeys = append(evictKeys, key)
		}
	}
	for _, evictKey := range evictKeys {
		delete(dst, evictKey)
	}
}

/*
 * Map Safe
 */

// MapOk
// check if k in m, never failure if m == nil
func MapOk[Map ~map[K]V, K comparable, V any](m Map, k K) bool {
	if len(m) == 0 {
		return false
	}
	_, ok := m[k]
	return ok
}

// MapElem
// get elem by k from m, never failure if m == nil
func MapElem[Map ~map[K]V, K comparable, V any](m Map, k K) V {
	if len(m) == 0 {
		return zero[V]()
	}
	return m[k]
}

// MapElemOk
// mapaccess2
func MapElemOk[Map ~map[K]V, K comparable, V any](m Map, k K) (V, bool) {
	if len(m) == 0 {
		return zero[V](), false
	}
	v, ok := m[k]
	return v, ok
}

func MapCopy[Map ~map[K]V, K comparable, V any](dst, src Map) Map {
	if len(dst) == 0 {
		return maps.Clone(src)
	}
	maps.Copy(dst, src)
	return dst
}

func MapDelete[Map ~map[K]V, K comparable, V any](m Map, k ...K) {
	for i := range k {
		delete(m, k[i])
	}
}

func MapClear[Map ~map[K]V, K comparable, V any](m ...Map) {
	for i := range m {
		clear(m[i])
	}
}

func MapKeys[Map ~map[K]V, K comparable, V any](m Map) []K {
	keys, index := make([]K, len(m)), 0
	for k := range m {
		keys[index] = k
		index++
	}
	return keys
}

func MapValues[Map ~map[K]V, K comparable, V any](m Map) []V {
	values, index := make([]V, len(m)), 0
	for _, v := range m {
		values[index] = v
		index++
	}
	return values
}

func MapInsert[Map ~map[E]struct{}, E comparable](m Map, e ...E) Map {
	if len(m) == 0 {
		return Slice2Map[[]E, map[E]struct{}, E](e)
	}
	mapInsert(m, e...)
	return m
}

func Slice2Map[Slice ~[]E, Map ~map[E]struct{}, E comparable](s Slice) Map {
	m := make(Map, len(s))
	mapInsert[Map, E](m, s...)
	return m
}

func mapInsert[Map ~map[E]struct{}, E comparable](m Map, e ...E) {
	for i := range e {
		m[e[i]] = struct{}{}
	}
}

func MapKeyMap[Map ~map[K]V, KeyMap ~map[K]struct{}, K comparable, V any](m Map) KeyMap {
	keyMap := make(KeyMap, len(m))
	for key := range m {
		keyMap[key] = struct{}{}
	}
	return keyMap
}

func MapValueMap[Map ~map[K]V, ValueMap ~map[V]struct{}, K, V comparable](m Map) ValueMap {
	valueMap := make(ValueMap, len(m))
	for key := range m {
		valueMap[m[key]] = struct{}{}
	}
	return valueMap
}

func MapsMerge[Map ~map[K]V, K comparable, V any](m ...Map) Map {
	all := make(Map, len(m)*MagicMapCap)

	for i := range m {
		all = MapCopy(all, m[i])
	}

	return all
}

func Slice2MapWhen[Slice ~[]E, E any, K comparable, V any](
	s Slice, predict Predicate2[E], transform Transform[E, K, V]) map[K]V {
	return lo.SliceToMap(lo.Filter(s, predict), transform)
}
