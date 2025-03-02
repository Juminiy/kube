package gorm_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/encrypt"
	"gorm.io/gorm"
	"testing"
)

type AppUser struct {
	gorm.Model
	Username string `mt:"unique"`
	Password string `gorm:"->:false;<-;not null"`
	Mobile   string `mt:"unique"`
	Email    string `mt:"unique"`
	Desc     string `gorm:"not null"`
	Status   int    `gorm:"default:1"`
}

func TestCreateAppUser(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&AppUser{}))
	appUser := AppUser{
		Username: "manboda",
		Password: encrypt.Md5("some_password"),
		Email:    "mi@cn",
		Desc:     util.MagicStr,
	}
	Err(t, _txTenant().Create(&appUser).Error)
	t.Log(Enc(appUser))
}

func TestParseWithTable(t *testing.T) {
	appUser := map[string]any{
		"Username": "iamhajimi",
		"Password": encrypt.Md5("some_password"),
		"Mobile":   "13888888888",
		"Desc":     util.MagicStr,
	}
	Err(t, _txTenant().Table(`tbl_app_user`).Create(appUser).Error)
	t.Log(Enc(appUser))
}

func TestParseWithTableCreateList(t *testing.T) {
	userList := []map[string]any{
		{"Username": "iamhajimi", "Password": encrypt.Md5("some_password"), "Mobile": "13888888888", "Desc": util.MagicStr},
		{"Username": "iammanboda", "Password": encrypt.Md5("some_password"), "Mobile": "13999999999", "Desc": util.MagicStr},
		{"Username": "imayoutube", "Password": encrypt.Md5("some_password"), "Mobile": "13000000000", "Desc": util.MagicStr},
		{"Username": "iambilibili", "Password": encrypt.Md5("some_password"), "Mobile": "13111111111", "Desc": util.MagicStr},
	}
	Err(t, _txTenant().Table(`tbl_app_user`).Create(&userList).Error)
	t.Log(Enc(userList))
}

func TestParseWithTableUpdate(t *testing.T) {
	Err(t, _txTenant().Table(`tbl_app_user`).
		Where("id = ?", 3).
		Updates(map[string]any{
			"desc": 114514,
		}).Error)
}

func TestParseWithTableQuery(t *testing.T) {
	var userList []struct {
		Username string
		Password string `gorm:"->:false"`
		Desc     string
		Status   int
	}
	Err(t, _txTenant().Table(`tbl_app_user`).Find(&userList).Error)
	t.Log(Enc(userList))
}

func TestSelectOmit(t *testing.T) {
	var userList []AppUser
	Err(t, _txTenant().Find(&userList).Error)
	t.Log(Enc(userList))
}

func TestQueryModelEqDest(t *testing.T) {
	var appUserList []AppUser
	Err(t, _txTenant().Find(&appUserList).Error)

	Err(t, _txTenant().Model(&AppUser{}).Find(&appUserList).Error)

	Err(t, _txTenant().Table(`tbl_app_user`).Find(&appUserList).Error)
}

type appUserStruct struct { // named
	Username string `mt:"unique"`
	Mobile   string `mt:"unique"`
	Email    string `mt:"unique"`
	Desc     string `gorm:"not null"`
}

func TestQueryModelNeqDest(t *testing.T) {
	// Model->SchemaStruct, Dest->named
	var appUser2 appUserStruct
	Err(t, _txTenant().Model(&AppUser{}).Find(&appUser2).Error)

	// Model->SchemaTable, Dest->named
	Err(t, _txTenant().Table(`tbl_app_user`).Find(&appUser2).Error)

	var appUser struct { // unnamed
		Username string `mt:"unique"`
		Mobile   string `mt:"unique"`
		Email    string `mt:"unique"`
		Desc     string `gorm:"not null"`
	}
	// Model->SchemaStruct, Dest->unnamed
	Err(t, _txTenant().Model(&AppUser{}).Find(&appUser).Error)

	// Model->SchemaTable, Dest->unnamed
	Err(t, _txTenant().Table(`tbl_app_user`).Find(&appUser).Error)
}

type appUserList []struct { // named
	Username string `mt:"unique"`
	Mobile   string `mt:"unique"`
	Email    string `mt:"unique"`
	Desc     string `gorm:"not null"`
}

func TestQueryModelNeqDestList(t *testing.T) {
	var list2 appUserList
	// Model->SchemaStruct, Dest->[]named
	Err(t, _txTenant().Model(&AppUser{}).Find(&list2).Error)

	// Model->SchemaTable, Dest->[]named
	Err(t, _txTenant().Table(`tbl_app_user`).Find(&list2).Error)

	var list []struct { // unnamed
		Username string `mt:"unique"`
		Mobile   string `mt:"unique"`
		Email    string `mt:"unique"`
		Desc     string `gorm:"not null"`
	}
	// Model->SchemaStruct, Dest->[]unnamed
	Err(t, _txTenant().Model(&AppUser{}).Find(&list).Error)

	// Model->SchemaTable, Dest->[]unnamed
	Err(t, _txTenant().Table(`tbl_app_user`).Find(&list).Error)
}

func txMixed() *gorm.DB {
	return _txTenant().Set("user_id", "3")
}
