package util

import (
	"encoding/json"
	"fmt"
	"html/template"
)

/*
 * Type Assert
 */

func AssertInt(v any) bool {
	switch v.(type) {
	case int, int64, int32, int16, int8:
		return true
	default:
		return false
	}
}

func AssertUint(v any) bool {
	switch v.(type) {
	case uint, uint64, uint32, uint16, uint8:
		return true
	default:
		return false
	}
}

func AssertFloat(v any) bool {
	switch v.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}

func AssertNumeric(v any) bool {
	return AssertInt(v) || AssertUint(v) || AssertFloat(v)
}

func AssertString(v any) bool {
	switch v.(type) {
	case string, []byte:
		return true
	default:
		return false
	}
}

func AssertStringLike(v any) bool {
	switch v.(type) {
	case fmt.Stringer:
		return true
	default:
		return AssertString(v)
	}
}

/*
 * Value Assert
 */

func AssertZero(v any) bool {
	switch z := v.(type) {
	case bool:
		return z // bool value exists conflict means in different case
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32:
		return z == 0
	case string:
		return len(z) == 0
	case []byte:
		return len(z) == 0
	case json.Number:
		return len(z) == 0
	case template.HTML:
		return len(z) == 0
	case template.URL:
		return len(z) == 0
	case template.JS:
		return len(z) == 0
	case template.CSS:
		return len(z) == 0
	case template.HTMLAttr:
		return len(z) == 0
	case nil:
		return true
	case fmt.Stringer:
		return len(z.String()) == 0
	case error:
		return z == nil || len(z.Error()) == 0
	default:
		return false
	}
}
