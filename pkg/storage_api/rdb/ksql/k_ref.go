package ksql

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast/safe_parse"
)

var _SP = safe_parse.Parse

const (
	Null = safe_parse.Category(0)
	Bool = safe_parse.Category(1)
	Num  = safe_parse.Category(2)
	Time = safe_parse.Category(3)
	Text = safe_parse.Category(4)
)

var _B2S = util.Bytes2StringNoCopy
var _S2B = util.String2BytesNoCopy
