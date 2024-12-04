package util

import (
	"archive/tar"
	"io"
	"time"
)

func TarIOReader2File(ioReader io.Reader, filePath string) (err error) {
	filePtr, err := OSOpenFileWithCreate(filePath)
	defer SilentCloseIO("tar file ptr", filePtr)
	if err != nil {
		return
	}

	tarFileWriter := tar.NewWriter(filePtr)
	defer SilentCloseIO("tar file writer", tarFileWriter)

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

func TarIOReaderToFileV2(ioReader io.Reader, filePath string) (err error) {
	filePtr, err := OSOpenFileWithCreate(filePath)
	defer SilentCloseIO("tar file ptr", filePtr)
	if err != nil {
		return
	}

	tarWriter := tar.NewWriter(filePtr)
	defer SilentCloseIO("tar file writer", tarWriter)
	_, err = io.Copy(tarWriter, ioReader)
	return
}
