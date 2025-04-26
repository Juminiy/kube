package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"strconv"
	"strings"
)

type IRReverse struct {
	Algo uint8
}

func (IRReverse) WithStringSalt(val string, salt ...string) string {
	return sha256Encrypt(val, util.StringConcat(salt...))
}

func (IRReverse) WithIntSalt(val string, salt ...int) string {
	str := strings.Builder{}
	str.WriteString(val)
	for _, saltInt := range salt {
		str.WriteString(strconv.Itoa(saltInt))
	}
	return sha256Encrypt(val, str.String())
}

func (IRReverse) WithUIntSalt(val string, salt ...uint) string {
	str := strings.Builder{}
	str.WriteString(val)
	for _, saltUint := range salt {
		str.WriteString(util.U64toa(uint64(saltUint)))
	}
	return sha256Encrypt(val, str.String())
}

func GetWithSalt(val string) (string, string) {
	salt := random.PasswordString(len(val))
	return sha256Encrypt(val, salt), salt
}

func sha256Encrypt(val, salt string) string {
	valWithSalt := util.StringConcat(val, salt)
	sha256Inst := sha256.New()
	_, err := sha256Inst.Write(util.String2BytesNoCopy(valWithSalt))
	if err != nil {
		stdlog.ErrorF("get encrypt sha256 val: %s with salt: %s error: %s", val, salt, err.Error())
		return valWithSalt
	}
	return hex.EncodeToString(sha256Inst.Sum(nil))
}
