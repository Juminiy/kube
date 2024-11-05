// Package mockv2 was generated
package mockv2

import (
	"github.com/brianvoe/gofakeit/v7"
)

type BoolFunc func() bool

var boolFunc = map[string]BoolFunc{
	defaultKey: defaultBool,
}

var defaultBool = gofakeit.Bool

var boolRule = rule{
	"bool:true":  true,
	"bool:false": false,
}
