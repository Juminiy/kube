package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
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

// Deprecated
func getObjectAbsPath(bucketName, objectPath string) string {
	if strings.HasPrefix(objectPath, "/") {
		return util.StringConcat(bucketName, objectPath)
	}
	return util.StringConcat(bucketName, "/", objectPath)
}
