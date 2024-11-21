package docker_file

// Expose
// EXPOSE <port> [<port>/<protocol>...]
type Expose struct {
	Port     int
	Protocol string
}
