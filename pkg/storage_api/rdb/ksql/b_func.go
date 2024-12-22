package ksql

type Brace struct {
	S, E string
	Do   BraceDo
}

func NewBrace[T ~string | rune | byte](S ...T) Brace {
	var st, ed string
	switch len(S) {
	case 1:
		st, ed = string(S[0]), string(S[0])
	case 2:
		st, ed = string(S[0]), string(S[1])
	}
	return Brace{
		S: st, E: ed, Do: func(str string) string {
			return st + str + ed
		},
	}
}

func (b Brace) Valid() bool {
	return b.Do != nil
}

func (b Brace) NotValid() bool {
	return b.Do == nil
}

type BraceDo func(str string) string

var accent = NewBrace("`")

var sQuote = NewBrace(`'`)

var quote = NewBrace(`"`)

var percent = NewBrace(`%`)

var bracket = NewBrace(`(`, `)`)

var none = NewBrace("")
