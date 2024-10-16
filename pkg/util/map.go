package util

import "maps"

const (
	MagicMapCap = (iota + 1) << 4
)

// MapOk
// check if k in m
func MapOk[Map ~map[K]V, K comparable, V any](m Map, k K) bool {
	if m == nil {
		return false
	}

	_, ok := m[k]
	return ok
}

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
