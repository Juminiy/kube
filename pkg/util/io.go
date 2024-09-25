package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"io"
)

func IOGetStr(readCloser io.ReadCloser) string {
	readBytes := make([]byte, 0, 4*Ki)
	_, err := readCloser.Read(readBytes)
	if err != nil {
		stdlog.Error(err)
	}
	return Bytes2StringNoCopy(readBytes)
}
