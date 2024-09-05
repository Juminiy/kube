//go:build windows

package util

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"os"
	"strings"
)

const (
	slash  = "\"
	slash2 = "\\"
)

func OSCreateAbsoluteDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

//OSCreateAbsolutePath
// A:\ A:\\
// B:\ B:\\
// C:\ C:\\
// D:\ D:\\
// WARNING: not test
func OSCreateAbsolutePath(filePath string) error {
	filePathAllowBySep := func(sep string) bool {
		parts := strings.Split(filePath, ":"+sep)
		return len(parts) != 0 &&
			len(parts[0]) == 1 &&
			strings.ToUpper(parts[0]) == parts[0]
	}
	if !filePathAllowBySep(slash) &&
		!filePathAllowBySep(slash2) {
		return fmt.Errorf("error filename: %s, filename must be an absolute path", filePath)
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
		stdlog.ErrorF("check file: %s exists error: %s", filePath, err.Error())
	}

	return false
}
