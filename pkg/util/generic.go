package util

func zero[V any]() V {
	var v V
	return v
}

func Zero[V any]() V {
	return zero[V]()
}
