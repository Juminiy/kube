//go:build windows

package internal_api

import (
	"strings"
)

const (
	Slash  = slash
	slash  = "\\"
	slash2 = "\\"
)

func IsAbsolutePath(filePath string) error {
	filePathAllowBySep := func(sep string) bool {
		parts := strings.Split(filePath, ":"+sep)
		return len(parts) != 0 &&
			len(parts[0]) == 1 &&
			strings.ToUpper(parts[0]) == parts[0]
	}
	if !filePathAllowBySep(slash) &&
		!filePathAllowBySep(slash2) {
		return notAbsPathErr(filePath)
	}
	return nil
}
