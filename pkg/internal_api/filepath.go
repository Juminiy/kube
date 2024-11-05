package internal_api

import (
	"fmt"
	//"github.com/Juminiy/kube/pkg/util"
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
	return err == nil || os.IsExist(err)
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
	return err == nil || os.IsExist(err)
}

func FileNotExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return os.IsNotExist(err)
}

//func CreateFile(filePath string) error {
//	filePtr, err := os.Create(filePath)
//	defer filePtr.Close()
//	return err
//}

func DeleteFile(filePath string) error {
	if FileExist(filePath) {
		return os.Remove(filePath)
	}
	return nil
}

func AppendFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, FilePerm)
}

func OverwriteFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, FilePerm)
}

//func AppendFileIfExist(filePath string) (*os.File, error) {
//	dir, _ := splitDirAndFileName(filePath)
//	if DirExist(dir) && FileExist(filePath) {
//		return AppendFile(filePath)
//	}
//	return nil, os.ErrNotExist
//}

// AppendCreateFile
// append file if exist and create file if not exist
func AppendCreateFile(filePath string) (*os.File, error) {
	dir, _ := splitDirAndFileName(filePath)
	if DirNotExist(dir) {
		err := CreateDir(dir)
		if err != nil {
			return nil, err
		}
	}
	if FileNotExist(filePath) {
		return os.Create(filePath)
	}
	return AppendFile(filePath)
}

// OverwriteCreateFile
// overwrite file if exist and create file if not exist
func OverwriteCreateFile(filePath string) (*os.File, error) {
	dir, _ := splitDirAndFileName(filePath)
	if DirNotExist(dir) {
		err := CreateDir(dir)
		if err != nil {
			return nil, err
		}
	}
	if FileNotExist(filePath) {
		return os.Create(filePath)
	}
	return OverwriteFile(filePath)
}

// SplitDirAndFileName
// +example
// "a" -> "", "a"
// "a/b" -> "a", "b"
// "a/b/c.d" -> "a/b","c.d"
// "/a/b/c.d" -> "/a/b","c.d"
func splitDirAndFileName(filePath string) (string, string) {
	lastSlashIndex := strings.LastIndex(filePath, slash)
	if lastSlashIndex == -1 {
		return "", filePath
	}
	return filePath[:lastSlashIndex], // dir
		filePath[lastSlashIndex+1:] // fileName
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

func notAbsPathErr(pathOrDir string) error {
	return fmt.Errorf("error path: %s, path must be an absolute path", pathOrDir)
}

func GetDirFileNames(dir string) ([]string, error) {
	entrys, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(entrys))
	for _, entry := range entrys {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}
	return names, nil
}
