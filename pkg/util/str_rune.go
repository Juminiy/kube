package util

import (
	"unicode"
)

// space run as string, space character as string
const (
	Tab     = "\t" // ASCII 9
	Newline = "\n" // ASCII 10
)

func IsSpace(s string) bool {
	for i := range s {
		if !isSpace(s[i]) {
			return false
		}
	}
	return true
}

func isSpace(b byte) bool {
	return unicode.IsSpace(rune(b))
}

// Printable and Showable ASCII Table
// referred from:
// https://en.wikipedia.org/wiki/ASCII
// https://www.ascii-code.com/
const (
	// From Null, following is ASCII 0 ~ 127
	Null = `\0` // ASCII 0

	Bell           = `\a` // ASCII 7
	Backspace      = `\b` // ASCII 8
	HorizontalTab  = `\t` // ASCII 9
	LineFeed       = `\n` // ASCII 10
	VerticalTab    = `\v` // ASCII 11
	FormFeed       = `\f` // ASCII 12
	CarriageReturn = `\r` // ASCII 13

	Escape = `\e` // ASCII 27

	Space            = ` ` // ASCII 32
	ExclamationMark  = `!` // ASCII 33
	DoubleQuotes     = `"` // ASCII 34
	SpeechMarks      = `"` // ASCII 34
	NumberSign       = `#` // ASCII 35
	Dollar           = `$` // ASCII 36
	PercentSign      = `%` // ASCII 37
	Ampersand        = `&` // ASCII 38
	SingleQuote      = `'` // ASCII 39
	OpenParenthesis  = `(` // ASCII 40
	OpenBracket      = `(` // ASCII 40
	CloseParenthesis = `)` // ASCII 41
	CloseBracket     = `)` // ASCII 41
	Asterisk         = `*` // ASCII 42
	Plus             = `+` // ASCII 43
	Comma            = `,` // ASCII 44
	HyphenMinus      = `-` // ASCII 45
	Period           = `.` // ASCII 46
	Dot              = `.` // ASCII 46
	FullStop         = `.` // ASCII 46
	Slash            = `/` // ASCII 47
	Divide           = `/` // ASCII 47

	// Zero~Nine = 0~9 // ASCII 48~57

	Colon              = `:` // ASCII 58
	Semicolon          = `;` // ASCII 59
	LessThan           = `<` // ASCII 60
	OpenAngledBracket  = `<` // ASCII 60
	Equals             = `=` // ASCII 61
	GreaterThan        = `>` // ASCII 62
	CloseAngledBracket = `>` // ASCII 62
	QuestionMark       = `?` // ASCII 63
	AtSign             = `@` // ASCII 64

	// Uppercase A~Z // ASCII 65~90

	OpeningBracket = `[` // ASCII 91
	Backslash      = `\` // ASCII 92
	ClosingBracket = `]` // ASCII 93
	Caret          = `^` // ASCII 94
	Circumflex     = `^` // ASCII 94
	Underscore     = `_` // ASCII 95
	GraveAccent    = "`" // ASCII 96

	// Lowercase a-z // ASCII 97~122

	OpeningBrace    = `{`   // ASCII 123
	VerticalBar     = `|`   // ASCII 124
	ClosingBrace    = `}`   // ASCII 125
	EquivalencySign = `~`   // ASCII 126
	Tilde           = `~`   // ASCII 126
	Delete          = `DEL` // ASCII 127
)
