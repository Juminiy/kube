package docker_api

const (
	hostURL       = "tcp://192.168.31.242:2375"
	clientVersion = "1.43"
)

var (
	testNewClient, testDockerClientError = New(hostURL, clientVersion)
)
