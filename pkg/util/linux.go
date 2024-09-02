package util

import (
	"fmt"
	"os"
	"strings"
)

const (
	slash = "/"
)

func OSCreateAbsoluteDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func OSCreateAbsolutePath(fileName string) error {
	if !strings.HasPrefix(fileName, slash) {
		return fmt.Errorf("error filename: %s filename must be an absolute path", fileName)
	}
	err := os.MkdirAll(fileName[:strings.LastIndex(fileName, slash)], os.ModePerm)
	if err != nil {
		return err
	}
	_, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	return err
}
