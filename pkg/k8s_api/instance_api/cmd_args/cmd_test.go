package cmd_args

import (
	"testing"
)

const test = "test"

func TestArgConcat(t *testing.T) {
	t.Log(LinuxHostNameCtl(test))
	t.Log(LinuxAddUser(test))
	t.Log(LinuxSetUserPassword(test))
	t.Log(LinuxServiceStart(test))
	t.Log(LinuxTouch(test))
	t.Log(LinuxEcho(test))
	t.Log(LinuxChmod("777", test))
	t.Log(LinuxMkdir(test))
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
	t.Log(s3fsM.Args())
}
