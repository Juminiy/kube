package cmd_args

import (
	"strings"
)

// Immutable Variables
// WARNING:
// 1. all single args must end with andAndNextLine
// 2. all group args must end with LinuxEchoCmdArgsFinished

// Linux Cmd
var (
	LinuxTerminalAlwaysOpen   = "tail -f /dev/null"
	LinuxPermitSSHLoginByRoot = "sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config"
)

// Ubuntu Cmd
var (
	UbuntuUpdateUpgrade   = "apt-get -y update && apt-get -y upgrade"
	UbuntuServiceSSHStart = "service ssh start"
)

// CentOS Cmd
var (
	CentOSUpdateUpgrade   = "yum -y update && yum -y upgrade"
	CentOSServiceSSHStart = "service sshd start"
)

var (
	andAndNextLine = " && \\"
	oneSpace       = " "
)

func ArgsConcat(s ...string) string {
	return strings.Join(s, andAndNextLine)
}

func ArgConcat(s ...string) string {
	return strings.Join(s, oneSpace)
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

func UbuntuInstall(software string) string {
	return ArgConcat("apt-get", "install", "-y", software)
}

func CentOSInstall(software string) string {
	return ArgConcat("yum", "install", "-y", software)
}
