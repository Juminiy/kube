package util

import (
	"errors"
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	syncerr "github.com/aws/smithy-go/sync"
	valyalabuffer "github.com/valyala/bytebufferpool"
	"io"
	"sync"
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
			stdlog.ErrorF("%s: %#v close error: %s", msg, closer, err.Error())
		}
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

/*func MustDetail(err error) {
	if err != nil {
		panic(err.Error())
	}
}*/

func MergeError(err ...error) error {
	return mergeErrorSep(", ", err...)
}

func mergeErrorSep(sep string, err ...error) error {
	if len(err) == 1 {
		return err[0]
	}
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

// ErrHandle
// goroutine safe
type ErrHandle struct {
	err    *syncerr.OnceErr
	errs   []error
	errsRw sync.RWMutex
}

func NewErrHandle() *ErrHandle {
	return &ErrHandle{
		err:  syncerr.NewOnceErr(),
		errs: make([]error, 0, MagicSliceCap),
	}
}

func (h *ErrHandle) Has(err ...error) bool {
	has := h.err.Err() != nil
	h.errsRw.Lock()
	defer h.errsRw.Unlock()
	for i := range err {
		if err[i] != nil {
			has = true
			h.err.SetError(err[i])
			h.errs = append(h.errs, err[i])
		}
	}
	return has
}

func (h *ErrHandle) HasStr(s ...string) bool {
	has := h.err.Err() != nil
	h.errsRw.Lock()
	defer h.errsRw.Unlock()
	for i := range s {
		if len(s[i]) > 0 {
			has = true
			newErr := errors.New(s[i])
			h.err.SetError(newErr)
			h.errs = append(h.errs, newErr)
		}
	}
	return has
}

func (h *ErrHandle) First() error {
	return h.err.Err()
}

func (h *ErrHandle) All(sep ...string) error {
	h.errsRw.RLock()
	defer h.errsRw.RUnlock()
	if len(sep) > 0 && len(sep[0]) > 0 {
		return mergeErrorSep(sep[0], h.errs...)
	}
	return mergeErrorSep("\n", h.errs...)
}

func (h *ErrHandle) AllStr(sep ...string) string {
	if err := h.All(sep...); err != nil {
		return err.Error()
	}
	return ""
}

var ErrFaked = errors.New("faked error")
