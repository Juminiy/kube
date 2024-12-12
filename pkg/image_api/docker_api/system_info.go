package docker_api

import (
	"github.com/Juminiy/kube/pkg/util/safe_json"
	dockersystem "github.com/docker/docker/api/types/system"
)

func (c *Client) SystemInfo() (dockersystem.Info, error) {
	return c.cli.Info(c.ctx)
}

func EncE(v any) ([]byte, error) {
	return safe_json.Goccy().Marshal(v)
}
func DecE(b []byte, v any) error {
	return safe_json.Goccy().Unmarshal(b, v)
}
