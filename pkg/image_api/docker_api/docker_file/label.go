package docker_file

// Label
// LABEL <key>=<value> [<key>=<value>...]
type Label struct {
	Kv  KeyEqVal
	Kvs []KeyEqVal
}
