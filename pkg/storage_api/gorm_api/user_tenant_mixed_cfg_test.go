package gorm_api

import (
	"encoding/json"
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

type Name string

func (n Name) MarshalJSON() ([]byte, error) {
	if n == "" || n == "null" {
		return nil, nil
	}
	return util.String2BytesNoCopy(string(n)), nil
}

func (n *Name) UnmarshalJSON(b []byte) error {
	if str := util.Bytes2StringNoCopy(b); str != "null" {
		*n = Name(str)
	}
	return nil
}

func TestMagicType(t *testing.T) {
	n := Name("null")
	bs, err := json.Marshal(n)
	Err(t, err)
	t.Logf("%s", bs)
}
