package docker_file

// Run
// +format(Shell form): RUN [OPTIONS] <command> ...
// +format(Exec form): RUN [OPTIONS] [ "<command>", ... ]
// +example: RUN source $HOME/.bashrc && echo $HOME
type Run struct {
	Exe
	Options []RunOption
}

type RunOption string

const (
	Mount    RunOption = "--mount"    // Minimum Dockerfile version 1.2
	Network  RunOption = "--network"  // Minimum Dockerfile version 1.3
	Security RunOption = "--security" // Minimum Dockerfile version 1.1.2-labs
)

// MountType
// RUN --mount=[type=<TYPE>][,option=<value>[,option=<value>]...]
type MountType string

const (
	Bind   MountType = "bind"
	Cache  MountType = "cache"
	Tmpfs  MountType = "tmpfs"
	Secret MountType = "secret"
	SSH    MountType = "ssh"
)

// NetworkType
// RUN --network=<TYPE>
type NetworkType string

const (
	Default NetworkType = "default"
	None    NetworkType = "none"
	Host    NetworkType = "host"
)

// SecurityType
// RUN --security=<sandbox|insecure>
type SecurityType string

const (
	Sandbox  SecurityType = "sandbox"
	Insecure SecurityType = "insecure"
)
