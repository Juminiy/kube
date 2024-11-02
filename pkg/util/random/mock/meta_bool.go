package mock

import "github.com/brianvoe/gofakeit/v7"

type BoolFunc func() bool

var boolFunc = map[string]BoolFunc{
	defaultKey: defaultBool,
}

var defaultBool = gofakeit.Bool

const (
	boolDefault = defaultKey
)

var boolRule = rule{
	"bool:true":  true,
	"bool:false": false,
}
