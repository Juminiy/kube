package docker_file

// From
// FROM [--platform=<platform>] <image> [AS <name>]
// FROM [--platform=<platform>] <image>[:<tag>] [AS <name>]
// FROM [--platform=<platform>] <image>[@<digest>] [AS <name>]
type From struct {
	Platform *string
	Image    Image
	AsName   *string
}

// Arg
// ARG  CODE_VERSION=latest

type Image struct {
	Repository string
	Tag        *string
	Digest     *string
}

type KeyEqVal struct {
	Key string
	Val string
}
