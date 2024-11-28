package util

import (
	valyalabuffer "github.com/valyala/bytebufferpool"
	"strings"
	"unicode"
)

// Snake2Camel
// +example: snake -> snake
// +example: snake_to_camel -> snakeToCamel
func Snake2Camel(snakeCase string) string {
	var camelCase string
	DoWithBuffer(func(buf *valyalabuffer.ByteBuffer) {
		snakeParts := strings.Split(snakeCase, "_")
		if len(snakeParts) == 0 {
			return
		}
		buf.WriteString(snakeParts[0]) // ignore first one
		for i := 1; i < len(snakeParts); i++ {
			_, _ = buf.WriteString(CapitalizeFirst(snakeParts[i]))
		}
		camelCase = buf.String()
	})
	return camelCase
}

// Camel2Snake
// +example: snake -> snake
// +example: snakeToCamel -> snake_to_camel
func Camel2Snake(camelCase string) string {
	var snakeCase string
	DoWithBuffer(func(buf *valyalabuffer.ByteBuffer) {
		for i, r := range camelCase {
			if unicode.IsUpper(r) {
				if i != 0 {
					_ = buf.WriteByte('_')
				}
				_, _ = buf.WriteString(string(unicode.ToLower(r)))
			} else {
				_ = buf.WriteByte(camelCase[i])
			}
		}

		snakeCase = buf.String()
	})
	return snakeCase
}

func CapitalizeFirst(str string) string {
	if len(str) > 0 {
		var follow string
		if len(str) > 1 {
			follow = str[1:]
		}
		return string(unicode.ToUpper(rune(str[0]))) + follow
	}
	return ""
}

// Camel2SnakeV2
// ArchOf -> arch_of
// GPUType -> gpu_type
// cpuCOUNT -> cpu_count
func Camel2SnakeV2(s string) string {
	return s
}
