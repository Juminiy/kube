package util

import "github.com/fatih/color"

func RedF(format string, v ...any) string {
	return color.RedString(format, v...)
}

func RedAny(v any) string {
	return color.RedString("%v", v)
}

func GreenF(format string, v ...any) string {
	return color.GreenString(format, v...)
}

func GreenAny(v any) string {
	return color.GreenString("%v", v)
}

// YN
// return Y true or N false
func YN(v bool) string {
	if v {
		return color.GreenString("Y")
	}
	return color.RedString("N")
}
