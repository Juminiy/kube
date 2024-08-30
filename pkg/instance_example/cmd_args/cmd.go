package cmd_args

import (
	"github.com/Juminiy/kube/pkg/util"
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
	linuxFileEtcS3cred = "/etc/s3cred"
)

type S3fsMount struct {
	Key        string
	Dir        string
	BucketName string
	MinioAddr  string
}

func (s *S3fsMount) Args() string {
	return ArgsConcat(
		// echo [AccessKeyID:SecretAccessKey] > /etc/s3cred
		LinuxEcho(s.Key, createOrOverwrite, linuxFileEtcS3cred),
		// chmod 600 /etc/s3cred
		LinuxChmod("600", linuxFileEtcS3cred),
		// mkdir -p [MountDir]
		LinuxMkdir(s.Dir),
		// s3fs [BucketName] [MountDir] -o passwd_file=/etc/s3cred,use_path_request_style,url=http://192.168.31.66:9000
		ArgConcat("s3fs", s.BucketName, s.Dir,
			"-o",
			util.StringConcat(
				"passwd_file=", linuxFileEtcS3cred,
				",use_path_request_style",
				",url=", util.URLWithHTTP(s.MinioAddr)),
		),
	)
}
