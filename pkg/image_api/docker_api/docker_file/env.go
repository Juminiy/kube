package docker_file

// Env
// ENV <key>=<value> [<key>=<value>...]
type Env struct {
	Kv  KeyEqVal
	Kvs []KeyEqVal
}
