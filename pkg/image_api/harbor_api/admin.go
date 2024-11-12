package harbor_api

import (
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/user"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

func (c *Client) CreateAdmin(userCreationReq *models.UserCreationReq) (*user.CreateUserCreated, error) {
	return c.AdminCreateUser(userCreationReq)
}

func (c *Client) DeleteAdmin(userID int64) (*user.DeleteUserOK, error) {
	return c.AdminDeleteUser(userID)
}

func (c *Client) AdminCreateUser(userCreationReq *models.UserCreationReq) (*user.CreateUserCreated, error) {
	return c.v2Cli.User.CreateUser(
		c.ctx,
		user.NewCreateUserParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithUserReq(userCreationReq),
	)
}

func (c *Client) AdminDeleteUser(userID int64) (*user.DeleteUserOK, error) {
	return c.v2Cli.User.DeleteUser(
		c.ctx,
		user.NewDeleteUserParams().
			WithContext(c.ctx).
			WithHTTPClient(c.httpCli).
			WithTimeout(c.httpTimeout).
			WithUserID(userID),
	)
}
