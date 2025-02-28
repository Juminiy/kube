package gorm_api

import "testing"

func TestClausePrimaryKey(t *testing.T) {
	var appUser AppUser
	appUser.ID = 3
	Err(t, _txTenant().First(&appUser).Error)
	t.Log(Enc(appUser))

	Err(t, _txTenant().Find(&appUser).Error)
	t.Log(Enc(appUser))
}
