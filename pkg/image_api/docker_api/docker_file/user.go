package docker_file

// User
// USER <user>[:<group>]
// USER <UID>[:<GID>]
type User struct {
	URep string
	GRep *string
}
