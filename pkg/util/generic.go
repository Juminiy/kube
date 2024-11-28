package util

func zero[V any]() V {
	var v V
	return v
}

func Zero[V any]() V {
	return zero[V]()
}

type Predicate[T any] func(item T) bool
type Predicate2[T any] func(item T, index int) bool

type Transform[T any, K comparable, V any] func(item T) (K, V)
type Transform2[T any, K comparable, V any] func(item T, index int) (K, V)
