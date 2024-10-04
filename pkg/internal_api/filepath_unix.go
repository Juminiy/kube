//go:build unix

package internal_api

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
