package util

import (
	"compress/gzip"
	"io"
)

func GzipIOReader2File(ioReader io.Reader, fileName string) (err error) {
	filePtr, err := OSOpenFileWithCreate(fileName)
	defer HandleCloseError("gzip file ptr", filePtr)
	if err != nil {
		return
	}

	gzipFileWriter := gzip.NewWriter(filePtr)
	defer HandleCloseError("gzip file writer", gzipFileWriter)

	_, err = io.Copy(gzipFileWriter, ioReader)
	if err != nil {
		return
	}

	err = gzipFileWriter.Flush()
	return
}
