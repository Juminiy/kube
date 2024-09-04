package random

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/teris-io/shortid"
)

func URLSafeString(size int) string {
	urlSafeString, err := shortid.Generate()
	if err != nil {
		stdlog.ErrorF("generate url safe string error: %s", err.Error())
		return password(9)
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
