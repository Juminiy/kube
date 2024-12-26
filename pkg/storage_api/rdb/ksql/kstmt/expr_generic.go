package kstmt

type AtLeastOne[T any] struct {
	First  T
	Remain []T
}

type Pair[T any] struct {
	First, Second T
}

type Empty struct{}

var NoEmpty = &Empty{}
