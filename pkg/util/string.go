package util

import (
	"strings"
)

const (
	MagicStr = "Ciallo~"
)

// String2BytesNoCopy
// unsafe, long string panic tested
func String2BytesNoCopy(s string) []byte {
	//sh := (*[2]uintptr)(unsafe.Pointer(&s))
	//bh := [3]uintptr{sh[0], sh[1], sh[1]}
	//return *(*[]byte)(unsafe.Pointer(&bh))
	return s2b(s)
}

func str2bsSafe(s *string) []byte {
	return []byte(*s)
}

// Bytes2StringNoCopy
// unsafe, long string panic tested
func Bytes2StringNoCopy(bs []byte) string {
	//bh := (*[3]uintptr)(unsafe.Pointer(&bs))
	//sh := [2]uintptr{bh[0], bh[1]}
	//return *(*string)(unsafe.Pointer(&sh))
	return b2s(bs)
}

func bs2strSafe(bs []byte) string {
	return string(bs)
}

func StringConcat(s ...string) string {
	return strings.Join(s, "")
}

func StringJoin(sep string, s ...string) string { return strings.Join(s, sep) }

func StringReplaceAlls(s, to string, from ...string) string {
	for _, fromElem := range from {
		s = strings.ReplaceAll(s, fromElem, to)
	}
	return s
}

func StringQuote(s string) string {
	return StringConcat("\"", s, "\"")
}

func StringSQuote(s string) string {
	return StringConcat("'", s, "'")
}

func StringPrefixIn(s string, p ...string) bool {
	for _, sp := range p {
		if strings.HasPrefix(s, sp) {
			return true
		}
	}
	return false
}

func StringDelete(s string, from ...string) string {
	for _, fe := range from {
		s = strings.ReplaceAll(s, fe, "")
	}
	return s
}

// a bit of wrong for type:rune or type:byte?
func StringSlice2RuneSlice(s []string) []rune {
	rs := make([]rune, len(s))
	for i := range s {
		if len(s[i]) > 0 {
			rs[i] = rune(s[i][0])
		}
	}
	return rs
}

func String2RuneSlice(s string) []rune {
	rs := make([]rune, len(s))
	for i := range s {
		rs[i] = rune(s[i])
	}
	return rs
}

func StringContainAny(s string, e ...string) bool {
	for i := range e {
		if strings.Contains(s, e[i]) {
			return true
		}
	}
	return false
}

func StringContainsAll(s string, e ...string) bool {
	for i := range e {
		if !strings.Contains(s, e[i]) {
			return false
		}
	}
	return true
}
