package harbor_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

func errorDetail(err error) error {
	payloadVal := safe_reflect.IndirectOf(err).StructFieldValue("Payload")
	switch errv := payloadVal.(type) {
	case *models.Error:
		detail, _ := errv.MarshalBinary()
		return errors.New(util.Bytes2StringNoCopy(detail))
	case *models.Errors:
		detail, _ := errv.MarshalBinary()
		return errors.New(util.Bytes2StringNoCopy(detail))
	default:
		stdlog.Warn("uncaught error type")
		return err
	}
}

func UnwrapErr[T any](v T, err error) (T, error) {
	if err != nil {
		return v, errorDetail(err)
	}
	return v, nil
}
