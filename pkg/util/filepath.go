package util

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"os"
)

func OSCreateAbsoluteDir(dir string) error {
	if err := internal_api.IsAbsolutePath(dir); err != nil {
		return err
	}
	return os.MkdirAll(dir, internal_api.DirPerm)
}

// OSCreateAbsolutePath
// unix
// /opt/xxx
// windows
// A:\ A:\\
// B:\ B:\\
// C:\ C:\\
// D:\ D:\\
func OSCreateAbsolutePath(filePath string) error {
	if err := internal_api.IsAbsolutePath(filePath); err != nil {
		return err
	}
	filePtr, err := internal_api.AppendCreateFile(filePath)
	defer SilentCloseIO("file ptr", filePtr)
	return err
}

func OSFilePathExists(filePath string) bool {
	return internal_api.FileExist(filePath)
}

func OSOpenFileWithCreate(filePath string) (*os.File, error) {
	return internal_api.AppendCreateFile(filePath)
}

func OSRemoveFile(filePath string) error {
	return internal_api.DeleteFile(filePath)
}
