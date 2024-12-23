package ksql

import "k8s.io/apimachinery/pkg/util/sets"

type Clause struct {
	Select   sets.Set[string]
	Distinct *bool
	Limit    *int
	Offset   int
	Order    *Order
	Group    *Group
}

type Order struct {
	Column string
	Desc   bool
}

type Group struct {
}
