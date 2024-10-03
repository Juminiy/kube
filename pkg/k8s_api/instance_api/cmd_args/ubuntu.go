package cmd_args

// Ubuntu Cmd
const (
	UbuntuUpdateUpgrade   = "apt-get -y update && apt-get -y upgrade"
	UbuntuServiceSSHStart = "service ssh start"
)

func UbuntuInstall(software string) string {
	return ArgConcat("apt-get", "install", "-y", software)
}
