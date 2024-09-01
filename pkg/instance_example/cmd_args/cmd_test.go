package cmd_args

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

var pfn = fmt.Println

const test = "test"

func TestArgConcat(t *testing.T) {
	_, err := pfn(LinuxHostNameCtl(test))
	_, err = pfn(LinuxAddUser(test))
	_, err = pfn(LinuxSetUserPassword(test))
	_, err = pfn(LinuxServiceStart(test))
	_, err = pfn(LinuxTouch(test))
	_, err = pfn(LinuxEcho(test))
	_, err = pfn(LinuxChmod("777", test))
	_, err = pfn(LinuxMkdir(test))

	util.SilentHandleError("", err)
}

func TestS3fsMount_Args(t *testing.T) {
	s3fsM := &S3fsMount{
		AccessKey:  "AccessKeyID:SecretAccessKey",
		MountDir:   "/mnt",
		BucketName: "s3fs-mount-bucket-test",
		MinioAddr:  "192.168.31.110:9000",
		S3CredNamingPolicy: S3CredNamingPolicy{
			GenMethod: S3CredNamingGenUUID,
		},
	}

	_, err := pfn(s3fsM.Args())
	util.SilentHandleError("", err)
}
