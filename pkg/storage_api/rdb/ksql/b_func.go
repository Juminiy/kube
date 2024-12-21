package ksql

type BraceF func(string) string

var accent = func(s string) string {
	return "`" + s + "`"
}

var sQuote = func(s string) string {
	return "'" + s + "'"
}

var quote = func(s string) string {
	return `"` + s + `"`
}

var percent = func(s string) string {
	return "%" + s + "%"
}

var bracket = func(s string) string {
	return "(" + s + ")"
}
