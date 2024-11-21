package docker_file

// Arg
// ARG <name>[=<default value>] [<name>[=<default value>]...]
type Arg struct {
	Kv  KeyEqVal
	Kvs KeyEqVal
}
