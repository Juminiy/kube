package cmd_args

import (
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

const (
	linuxDirEtc = "/etc"
)
