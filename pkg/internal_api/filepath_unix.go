//go:build unix

package internal_api

import "path/filepath"

const (
	Slash = slash
	slash = "/"
)

func IsAbsolutePath(filePath string) error {
	if !filepath.IsAbs(filePath) {
		return notAbsPathErr(filePath)
	}
	return nil
}
