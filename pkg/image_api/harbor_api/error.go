package harbor_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

func (c *Client) ErrorDetail(err error) string {
	if err == nil {
		return ""
	}
	err = errorDetail(err)
	if err != nil {
		return err.Error()
	}
	return ""
}

func errorDetail(err error) error {
	payloadVal := safe_reflect.IndirectOf(err).StructFieldValue("Payload")
	if payLoadErr, ok := payloadVal.(*models.Error); ok {
		bs, err := payLoadErr.MarshalBinary()
		if err != nil {
			return err
		}
		return errors.New(util.Bytes2StringNoCopy(bs))
	}
	return nil
}

func UnwrapErr[T any](v T, err error) (T, error) {
	if err != nil {
		return v, errorDetail(err)
	}
	return v, nil
}
