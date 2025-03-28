package random

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/zerobuf"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/teris-io/shortid"
	valyalabuffer "github.com/valyala/bytebufferpool"
	"math/rand/v2"
	"strings"
)

const (
	magicFix8  = 8
	magicFix16 = 16
	magicFix32 = 32
	magicFix64 = 64
)

func URLSafeString(size int) string {
	urlSafeString, err := shortid.Generate()
	if err != nil {
		stdlog.ErrorF("generate url safe string error: %s", err.Error())
		return AlphaNumericString(size)
	}
	return urlSafeString
}

func IDString(size int) string {
	return AlphaNumericString(size)
}

func ID() string {
	return AlphaNumericString(magicFix8)
}

func PasswordString(size int) string {
	return AlphaNumericString(size)
}

func Password() string {
	return AlphaNumericString(magicFix8)
}

func LowerCaseString(size int) string {
	return gofakeit.Password(true, false, false, false, false, size)
}

func UpperCaseString(size int) string {
	return gofakeit.Password(false, true, false, false, false, size)
}

func AlphaString(size int) string {
	return gofakeit.Password(true, true, false, false, false, size)
}

func NumericString(size int) string {
	return gofakeit.Password(false, false, true, false, false, size)
}

func AlphaNumericString(size int) string {
	return gofakeit.Password(true, true, true, false, false, size)
}

func SymbolString(size int) string {
	return gofakeit.Password(true, true, true, true, false, size)
}

func AllString(size int) string {
	return gofakeit.Password(true, true, true, true, true, size)
}

func FileNameString(ext string) string {
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return util.StringConcat(
		strings.TrimSpace(gofakeit.ProductName()),
		URLSafeString(rand.IntN(magicFix8)),
		ext,
	)
}

// Integer
// size = 1 -> 0~9
// size = 2 -> 10~99
// size = 3 -> 100~999
// ...
// size = n -> 10^(n-1) ~ 10^(n)-1
// or only byte + '0'
func Integer(size int) string {
	var integerStr string
	util.DoWithBuffer(func(buf *valyalabuffer.ByteBuffer) {
		for range size {
			_ = buf.WriteByte(byte(rand.IntN(10) + '0'))
		}
		integerStr = buf.String()
	})
	return integerStr
}

func NumericVerify(size int) string {
	return Integer(size)
}

const (
	alphaUStr  = "ABCDEFGHJKMNPQRSTUVWXYZ"
	alphaLStr  = "abcdefghjkmnpqrstuvwxyz"
	alphaStr   = alphaUStr + alphaLStr
	numericStr = "23456789"
)

func AlphaVerify(size int) string {
	return fromString(alphaStr, size)
}

func AlphaNumericVerify(size int) string {
	return fromString(numericStr, size)
}

func fromString(s string, size int) string {
	vbuf := zerobuf.Get()
	defer vbuf.Free()

	for range size {
		vbuf.WriteByte(s[rand.IntN(len(s))])
	}

	return vbuf.String()
}
