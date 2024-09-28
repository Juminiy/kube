package util

func MapOk[Map ~map[K]V, K comparable, V any](m Map, k K) bool {
	if m == nil {
		return false
	}

	_, ok := m[k]
	return ok
}
