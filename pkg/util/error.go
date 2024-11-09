package util

import (
	"errors"
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	valyalabuffer "github.com/valyala/bytebufferpool"
	"io"
)

// SilentPanic
// only used in dev env or test env, _test file
// not to use in production env
func SilentPanic(err error) {
	if err != nil {
		stdlog.Panic(err)
	}
}

// SilentFatal
// only used in dev env or test env, _test file
// not to use in production env
func SilentFatal(err error) {
	if err != nil {
		stdlog.Fatal(err)
	}
}

// SilentFatalf
// only used in dev env or test env, _test file
// not to use in production env
func SilentFatalf(detail string, err error) {
	if err != nil {
		stdlog.FatalF("%s: %s", detail, err.Error())
	}
}

// SilentError
// only used in dev env or test env, _test file
// not to use in production env
func SilentError(err error) {
	if err != nil {
		stdlog.Error(err.Error())
	}
}

// SilentErrorf
// only used in dev env or test env, _test file
// not to use in production env
func SilentErrorf(detail string, err error) {
	if err != nil {
		stdlog.ErrorF("%s: %s", detail, err.Error())
	}
}

// SilentCloseIO
// handle io closer error
func SilentCloseIO(msg string, closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			stdlog.ErrorF(msg+" instance: %#v close error: %s", closer, err.Error())
		}
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MergeError(err ...error) error {
	return mergeErrorSep(", ", err...)
}

func mergeErrorSep(sep string, err ...error) error {
	var (
		hasErr bool
		errStr string
	)
	DoWithBuffer(func(buf *valyalabuffer.ByteBuffer) {
		for i := range err {
			if err[i] != nil {
				_, _ = buf.WriteString(fmt.Sprintf("error[%d]: %s%s", i, err[i].Error(), sep))
				hasErr = true
			}
		}
		errStr = buf.String()
	})
	if hasErr {
		return errors.New(errStr)
	}
	return nil
}
