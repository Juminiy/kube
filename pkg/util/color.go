package util

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

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

type ColorValue struct {
	Color color.Attribute
	Value any
}

func Colorf(val ...ColorValue) {
	fmtStrBuilder := strings.Builder{}
	valSlice := make([]any, len(val))
	for i := range val {
		fmtStrBuilder.WriteString("%v")
		valSlice[i] = color.New(val[i].Color).SprintFunc()(val[i].Value)
	}
	fmtStrBuilder.WriteByte('\n')
	fmt.Printf(fmtStrBuilder.String(), valSlice...)
}
