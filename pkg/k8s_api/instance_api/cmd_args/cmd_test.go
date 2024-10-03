package cmd_args

import (
	"testing"

	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

const test = "test"

func TestArgConcat(t *testing.T) {
	stdlog.Info(LinuxHostNameCtl(test))
	stdlog.Info(LinuxAddUser(test))
	stdlog.Info(LinuxSetUserPassword(test))
	stdlog.Info(LinuxServiceStart(test))
	stdlog.Info(LinuxTouch(test))
	stdlog.Info(LinuxEcho(test))
	stdlog.Info(LinuxChmod("777", test))
	stdlog.Info(LinuxMkdir(test))

}

func TestS3fsMount_Args(t *testing.T) {
	s3fsM := &S3fsMount{
		AccessKey:  "AccessKeyID:SecretAccessKey",
		MountDir:   "/mnt",
		BucketName: "s3fs-mount-bucket-test",
		MinioAddr:  "minio.local",
		S3CredNamingPolicy: S3CredNamingPolicy{
			GenMethod: S3CredNamingGenUUID,
		},
	}

	stdlog.Info(s3fsM.Args())
}
