package cmd_args

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/google/uuid"
	"strings"
)

// Immutable Variables
// WARNING:
// 1. all single args must end with andAndNextLine
// 2. all group args must end with LinuxEchoCmdArgsFinished

// Linux Cmd
const (
	LinuxTerminalAlwaysOpen   = "tail -f /dev/null"
	LinuxPermitSSHLoginByRoot = "sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config"
)

// Ubuntu Cmd
const (
	UbuntuUpdateUpgrade   = "apt-get -y update && apt-get -y upgrade"
	UbuntuServiceSSHStart = "service ssh start"
)

// CentOS Cmd
const (
	CentOSUpdateUpgrade   = "yum -y update && yum -y upgrade"
	CentOSServiceSSHStart = "service sshd start"
)

// linux const var
const (
	andAnd            = " && "
	oneSpace          = " "
	createOrAppend    = ">>"
	createOrOverwrite = ">"
	fileHide          = "."
	dirSlash          = "/"
	underLine         = "_"
	linuxFileMaxLen   = 255
)

func ArgsConcat(s ...string) string {
	return strings.Join(s, andAnd)
}

func ArgConcat(s ...string) string {
	return strings.Join(s, oneSpace)
}

func ArgConcat2(s0 []string, s ...string) string {
	return ArgConcat(ArgConcat(s0...), ArgConcat(s...))
}

func ArgConcat3(s0, s1 []string, s ...string) string {
	return ArgConcat(ArgConcat(s0...), ArgConcat(s1...), ArgConcat(s...))
}

func LinuxHostNameCtl(hostname string) string {
	return ArgConcat("hostnamectl", "set-hostname", "--permanent", hostname)
}

func LinuxAddUser(username string) string {
	return ArgConcat("adduser", "-m", "-s", "/bin/bash", username)
}

func LinuxSetUserPassword(password string) string {
	return ArgConcat("echo", "-e", password+"\n"+password, "|", "passwd")
}

func LinuxServiceStart(service string) string {
	return ArgConcat("service", service, "start")
}

func LinuxTouch(filepath string) string {
	return ArgConcat("touch", filepath)
}

func LinuxEcho(s ...string) string {
	return ArgConcat2([]string{"echo"}, s...)
}

func LinuxChmod(pri, filepath string) string {
	return ArgConcat("chmod", pri, filepath)
}

func LinuxMkdir(dir string) string {
	return ArgConcat("mkdir", "-p", dir)
}

func UbuntuInstall(software string) string {
	return ArgConcat("apt-get", "install", "-y", software)
}

func CentOSInstall(software string) string {
	return ArgConcat("yum", "install", "-y", software)
}

const (
	linuxDirEtc        = "/etc"
	linuxFileEtcS3cred = "/etc/s3cred"
	linuxFileS3Cred    = "s3cred"
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
	S3CredNamingGenAdmin = iota
	S3CredNamingGenUUID
	S3CredNamingGenRandomString15
	S3CredNamingGenRandomString31
	S3CredNamingGenRandomString63
	S3CredNamingGenRandomString127
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
