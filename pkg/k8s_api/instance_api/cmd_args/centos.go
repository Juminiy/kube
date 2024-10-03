package cmd_args

// CentOS Cmd
const (
	CentOSUpdateUpgrade   = "yum -y update && yum -y upgrade"
	CentOSServiceSSHStart = "service sshd start"
)

func CentOSInstall(software string) string {
	return ArgConcat("yum", "install", "-y", software)
}
