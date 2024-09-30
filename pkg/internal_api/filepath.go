package internal_api

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	DirPerm  = 0755
	FilePerm = 0666
)

func DirExist(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsExist(err)
}

func DirNotExist(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsNotExist(err)
}

func CreateDir(dir string) error {
	if DirNotExist(dir) {
		return os.MkdirAll(dir, DirPerm)
	}
	return nil
}

func DeleteDir(dir string) error {
	if DirExist(dir) {
		return os.Remove(dir)
	}
	return nil
}

func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return os.IsExist(err)
}

func FileNotExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return os.IsNotExist(err)
}

func OpenFileIfExist(filePath string) (*os.File, error) {
	dir, _ := SplitDirAndFileName(filePath)
	if DirExist(dir) && FileExist(filePath) {
		return openFile(filePath)
	}
	return nil, os.ErrNotExist
}

func OpenFileWithCreateIfNotExist(filePath string) (*os.File, error) {
	dir, _ := SplitDirAndFileName(filePath)
	if DirNotExist(dir) {
		err := CreateDir(dir)
		if err != nil {
			return nil, err
		}
	}
	if FileNotExist(filePath) {
		return os.Create(filePath)
	}
	return openFile(filePath)
}

func DeleteFile(filePath string) error {
	if FileExist(filePath) {
		return os.Remove(filePath)
	}
	return nil
}

// SplitDirAndFileName
// +example
// "a" -> "", "a"
// "a/b" -> "a", "b"
// "a/b/c.d" -> "a/b","c.d"
// "/a/b/c.d" -> "/a/b","c.d"
func SplitDirAndFileName(filePath string) (string, string) {
	lastSlashIndex := strings.LastIndex(filePath, slash)
	if lastSlashIndex == -1 {
		return "", filePath
	}
	return filePath[:lastSlashIndex], // dir
		filePath[lastSlashIndex+1:] // fileName
}

func openFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, FilePerm)
}

func GetWorkPath(s ...string) (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filePathJoin(workDir, s...), nil
}

func filePathJoin(base string, s ...string) string {
	return filepath.Join(base, filepath.Join(s...))
}
