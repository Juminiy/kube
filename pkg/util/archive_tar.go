package util

import (
	"archive/tar"
	"io"
	"time"
)

func TarIOReader2File(ioReader io.Reader, filePath string) (err error) {
	filePtr, err := OSOpenFileWithCreate(filePath)
	defer HandleCloseError("tar file ptr", filePtr)
	if err != nil {
		return
	}

	tarFileWriter := tar.NewWriter(filePtr)
	defer HandleCloseError("tar file writer", tarFileWriter)

	timeNow := time.Now()
	err = tarFileWriter.WriteHeader(&tar.Header{
		Name:       filePath,
		Mode:       FileMaxPerm,
		ModTime:    timeNow,
		AccessTime: timeNow,
		ChangeTime: timeNow,
	})
	if err != nil {
		return
	}

	_, err = io.Copy(tarFileWriter, ioReader)
	return
}
