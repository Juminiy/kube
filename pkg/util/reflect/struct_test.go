package reflect

import (
	"testing"
)

func TestParseStructTag(t *testing.T) {
	type T0 struct {
		F0 string `gorm:"column:user_name;type:varchar(128);comment:user's name, account's name" json:"f0,omitempty" app:"name"`
		F1 int    `app:"i"`
	}

	tagMap := Of(T0{}).ParseStructTag("gorm")
	t.Log(tagMap)
	t.Log(tagMap.ParseGetTagValV("F0", "app"))
}
