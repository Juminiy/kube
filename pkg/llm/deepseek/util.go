package deepseek

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
)

var B2s = util.Bytes2StringNoCopy
var S2b = util.String2BytesNoCopy
var Enc = safe_json.GoCCY().Marshal
var Dec = safe_json.GoCCY().Unmarshal
var Pretty = safe_json.Pretty
