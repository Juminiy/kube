package kstmt

type AtLeastOne[T any] struct {
	First  T
	Remain []T
}

type Empty struct{}
