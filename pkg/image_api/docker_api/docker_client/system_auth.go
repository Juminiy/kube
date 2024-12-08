package docker_client

import (
	"github.com/docker/docker/api/types/registry"
)

func (c *Client) RegistryLogin(auth registry.AuthConfig) (authResp registry.AuthenticateOKBody, err error) {
	r := c.post("/auth")
	r.r.SetBody(auth)

	resp, err := r.do()
	if err != nil {
		return
	}
	err = DecE(resp.Body(), &authResp)
	return
}
