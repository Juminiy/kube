package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_go"
	"github.com/minio/minio-go/v7"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ObjectConfig struct {
	// +required
	BucketName string

	// +optional
	ObjectPath string

	// +required
	ObjectName string

	// +optional, only required in Client.PutObject
	ObjectSize int64
}

func (c *ObjectConfig) ObjectAbsPath() string {
	return c.BucketName + "/" + c.ObjectAbsName()
}

func (c *ObjectConfig) ObjectAbsName() string {
	if strings.HasPrefix(c.ObjectPath, "/") {
		c.ObjectPath = c.ObjectPath[1:]
	}
	suf, pref := strings.HasSuffix(c.ObjectPath, "/"),
		strings.HasPrefix(c.ObjectName, "/")
	if suf && pref {
		return c.ObjectPath + c.ObjectName[1:]
	} else if suf || pref {
		return c.ObjectPath + c.ObjectName
	}
	return c.ObjectPath + "/" + c.ObjectName
}

func (c *Client) TempGetObject(objectConfig *ObjectConfig, expiry time.Duration) (*url.URL, error) {
	reqParams := make(url.Values)
	reqParams.Set(
		"response-content-disposition",
		util.StringConcat("attachment; filename=", "\"", objectConfig.ObjectName, "\""),
	)

	presignedURL, err := c.mc.PresignedGetObject(
		c.ctx,
		objectConfig.BucketName,
		objectConfig.ObjectAbsName(),
		expiry,
		reqParams)
	if err != nil {
		stdlog.ErrorF("minio presigned get object: %s error: %s", objectConfig.ObjectAbsPath(), err.Error())
	}
	return presignedURL, err
}

func (c *Client) TempPutObject(objectConfig *ObjectConfig, expiry time.Duration) (*url.URL, error) {
	presignedURL, err := c.mc.PresignedPutObject(
		c.ctx,
		objectConfig.BucketName,
		objectConfig.ObjectAbsName(),
		expiry)
	if err != nil {
		stdlog.ErrorF("minio presigned put object: %s error: %s", objectConfig.ObjectAbsPath(), err.Error())
	}
	return presignedURL, err
}

func (c *Client) TempGetObjectList(objectConfigs []ObjectConfig, expiry time.Duration) ([]*url.URL, error) {
	tempURLs := make([]*url.URL, len(objectConfigs))
	fns := make([]util.Func, len(objectConfigs))
	for i, config := range objectConfigs {
		fns[i] = func() error {
			var tempErr error
			tempURLs[i], tempErr = c.TempGetObject(&config, expiry)
			return tempErr
		}
	}

	err := safe_go.DryRun(fns...)
	return tempURLs, err
}

func (c *Client) ObjectExists(objectConfig *ObjectConfig) (bool, error) {
	_, err := c.mc.StatObject(c.ctx, objectConfig.BucketName, objectConfig.ObjectAbsName(), minio.StatObjectOptions{})
	switch merr := err.(type) {
	case nil:
		return true, nil
	case minio.ErrorResponse:
		if merr.StatusCode == http.StatusNotFound {
			return false, nil
		} else {
			return false, err
		}
	default:
		return false, err
	}
}

func (c *Client) PutObject(objectConfig *ObjectConfig, input io.Reader) (minio.UploadInfo, error) {
	if objectConfig.ObjectSize == 0 {
		objectConfig.ObjectSize = -1
	}
	return c.mc.PutObject(c.ctx,
		objectConfig.BucketName,
		objectConfig.ObjectAbsName(),
		input,
		objectConfig.ObjectSize,
		minio.PutObjectOptions{})
}

// Deprecated
func getObjectAbsPath(bucketName, objectPath string) string {
	if strings.HasPrefix(objectPath, "/") {
		return util.StringConcat(bucketName, objectPath)
	}
	return util.StringConcat(bucketName, "/", objectPath)
}
