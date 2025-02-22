package compress

import (
	"bytes"
	"compress/bzip2"
	"github.com/Juminiy/kube/pkg/util"
	bz2 "github.com/dsnet/compress/bzip2"
	"io"
)

func Bzip2Encode(b []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	bWtr, err := bz2.NewWriter(buf, &bz2.WriterConfig{Level: bz2.BestCompression})
	if err != nil {
		return nil, err
	}
	_, err = bWtr.Write(b)
	if err != nil {
		return nil, err
	}
	util.SilentCloseIO("bzip2 writer", bWtr)
	return buf.Bytes(), nil
}

func Bzip2Decode(rdr io.Reader) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	bRdr, err := bz2.NewReader(rdr, &bz2.ReaderConfig{})
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buf, bRdr)
	if err != nil {
		return nil, err
	}
	util.SilentCloseIO("bzip2 reader", bRdr)
	return buf.Bytes(), err
}

// Bzip2DecodeStdlib
// Read nothing ??? why
func Bzip2DecodeStdlib(rdr io.Reader) ([]byte, error) {
	b := make([]byte, 0, util.Ki)
	_, err := bzip2.NewReader(rdr).Read(b)
	return b, err
}
