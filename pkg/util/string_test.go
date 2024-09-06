package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"testing"
)

func TestBytes2StringNoCopy(t *testing.T) {
	bytesOf := []byte{104, 98, 111}
	stdlog.Info(string(bytesOf))
	stdlog.Info(Bytes2StringNoCopy(bytesOf))
}

func TestString2BytesNoCopy(t *testing.T) {
	strOf := "Alan"
	stdlog.Info([]byte(strOf))
	stdlog.Info(String2BytesNoCopy(strOf))
}
