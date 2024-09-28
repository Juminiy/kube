package util

import (
	"unicode"
)

// space run as string, space character as string
const (
	Space          = " "  // ASCII 32
	Tab            = "\t" // ASCII 9
	Newline        = "\n" // ASCII 10
	CarriageReturn = "\r" // ASCII 13
	VerticalTab    = "\v" // ASCII 11
	FormFeed       = "\f" // ASCII 12
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
