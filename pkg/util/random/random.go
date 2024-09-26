package random

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/teris-io/shortid"
	"math/rand"
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
		return password(size)
	}
	return urlSafeString
}

var IDString = password
var PasswordString = password

func password(size int) string {
	return gofakeit.Password(
		true,
		true,
		true,
		false,
		false,
		size,
	)
}

func FileNameString(ext string) string {
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return util.StringConcat(
		strings.TrimSpace(gofakeit.ProductName()),
		URLSafeString(rand.Intn(magicFix8)),
		ext,
	)
}
