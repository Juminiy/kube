package cmd_args

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/google/uuid"
)

type (
	S3fsMount struct {
		// Minio AccessKey (short for [AccessKeyID:SecretAccessKey])
		AccessKey string

		// Mount Directory in Linux
		MountDir string

		// Mount BucketName in Minio Cluster
		BucketName string

		// MinioAddr (short for Minio Address)
		MinioAddr string

		// +optional
		S3CredNamingPolicy
	}
	S3CredNamingPolicy struct {
		// filename begin with .
		// for example: /etc/.s3_cred_xxx
		// default is NoHidden=false -> Hidden
		// +optional
		NoHidden bool

		// filename in linux max len is 255
		// +optional
		MaxLen int

		// name gen method
		// +optional
		GenMethod uint8
	}
)

const (
	S3CredNamingGenAdmin = iota + 1
	S3CredNamingGenUUID
	S3CredNamingGenRandomString15
	S3CredNamingGenRandomString31
	S3CredNamingGenRandomString63
	S3CredNamingGenRandomString127
)

const (
	linuxFileEtcS3cred = "/etc/s3cred"
	linuxFileS3Cred    = "s3cred"
)

func (s *S3fsMount) Args() string {
	s3CredPath := s.nameS3Cred()

	return ArgsConcat(
		// echo [AccessKeyID:SecretAccessKey] > /etc/s3cred
		LinuxEcho(s.AccessKey, createOrOverwrite, s3CredPath),
		// chmod 600 /etc/s3cred
		LinuxChmod("600", s3CredPath),
		// mkdir -p [MountDir]
		LinuxMkdir(s.MountDir),
		// s3fs [BucketName] [MountDir] -o passwd_file=/etc/s3cred,use_path_request_style,url=[MinioAddress]
		ArgConcat("s3fs", s.BucketName, s.MountDir,
			"-o",
			util.StringConcat(
				"passwd_file=", s3CredPath,
				",use_path_request_style",
				",url=", util.URLWithHTTP(s.MinioAddr)),
		),
	)
}

func (s *S3fsMount) nameS3Cred() string {
	var (
		filePath string
		hideFile = fileHide
	)
	if s.NoHidden {
		hideFile = ""
	}

	switch s.S3CredNamingPolicy.GenMethod {
	case S3CredNamingGenAdmin:
		filePath = util.StringConcat(
			linuxDirEtc,
			dirSlash,
			hideFile,
			linuxFileS3Cred,
		)
	case S3CredNamingGenUUID:
		filePath = util.StringConcat(
			linuxDirEtc,
			dirSlash,
			hideFile,
			linuxFileS3Cred,
			underLine,
			uuid.NewString(),
		)
	}

	if len(filePath) > linuxFileMaxLen {
		filePath = filePath[:linuxFileMaxLen-1]
	}
	if s.MaxLen > 0 && len(filePath) > s.MaxLen {
		filePath = filePath[:s.MaxLen-1]
	}
	return filePath
}
