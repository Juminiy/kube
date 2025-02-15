package encrypt

import (
	"crypto/md5"
	"github.com/Juminiy/kube/pkg/util"
)

func Md5(s string) string {
	return Md5Bytes2String(md5.Sum(s2b(s)))
}

func Md5Bytes2String(md5Arr [16]byte) string {
	return b2s(md5Arr[:16])
}

var s2b = util.String2BytesNoCopy
var b2s = util.Bytes2StringNoCopy
