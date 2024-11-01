package mock

type BoolFunc func() bool

type UintFunc func() uint64

type IntFunc func() int64

type FloatFunc func() float64

type StringFunc func() string

var boolFunc = map[string]BoolFunc{}

var uintFunc = map[string]UintFunc{}

var intFunc = map[string]IntFunc{}

var floatFunc = map[string]FloatFunc{}

var stringFunc = map[string]StringFunc{}
