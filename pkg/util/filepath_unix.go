//go:build unix

package util

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"os"
	"path/filepath"
	"strings"
)

const (
	slash = "/"
)

func OSCreateAbsoluteDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func OSCreateAbsolutePath(filePath string) error {
	if !filepath.IsAbs(filePath) {
		return fmt.Errorf("error filename: %s filename must be an absolute path", filePath)
	}
	err := os.MkdirAll(filePath[:strings.LastIndex(filePath, slash)], os.ModePerm)
	if err != nil {
		return err
	}
	_, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	return err
}

func OSFilePathExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		stdlog.InfoF("file %s does not exist", filePath)
	} else {
		stdlog.ErrorF("check file exists error: %s", err.Error())
	}

	return false
}

func OSOpenFileWithCreate(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
}

func OSRemoveFile(filePath string) error {
	if OSFilePathExists(filePath) {
		return os.Remove(filePath)
	}
	return nil
}
