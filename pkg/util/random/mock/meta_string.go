package mock

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
)

type StringFunc func() string

var stringFunc = map[string]StringFunc{}

var defaultString = func() string {
	return gofakeit.LetterN(stringDefaultMaxLen)
}

const (
	stringDefaultMinLen = 1
	stringDefaultMaxLen = 16
	stringDefaultRatio  = 2
	stringMaxLen        = util.Ki
)

var stringRule = rule{
	"string:char":    []rune{},
	"string:len:min": stringDefaultMinLen,
	"string:len:max": stringDefaultMaxLen,
}
