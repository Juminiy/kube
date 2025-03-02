package gorm_api

import (
	"testing"
)

/*// gorm do not support MapType Alias
type ConsumerMap map[string]any

func (m ConsumerMap) BeforeCreate(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		m["UserID"] = userID
	}
	return nil
}

func (m ConsumerMap) AfterCreate(tx *gorm.DB) error {
	delete(m, "UserID")
	return nil
}*/

func TestCallbacksBeforeCreateMapButUnsupported(t *testing.T) {
	// unsupported type: panic
	/*var consumerMapV2 = ConsumerMap{
		"AppID": 44,
	}
	Err(t, txMixed().Table(`tbl_consumer`).Create(&consumerMapV2).Error)
	// create ? user_id ?
	t.Log(Enc(consumerMapV2))*/
}

func TestCallbacksBeforeCreate(t *testing.T) {
	// create Struct Hook is success
	var consumerStruct = Consumer{
		AppID: 11,
	}
	Err(t, txMixed().Create(&consumerStruct).Error)
	// create with user_id
	t.Log(Enc(consumerStruct))

	// create Map
	var consumerMap = map[string]any{
		"AppID": 22,
	}
	Err(t, txMixed().Model(&Consumer{}).Create(&consumerMap).Error)
	// create with no user_id
	t.Log(Enc(consumerMap))

	var consumerMap2 = map[string]any{
		"AppID": 33,
	}
	Err(t, txMixed().Table(`tbl_consumer`).Create(&consumerMap2).Error)
	// create with no user_id
	t.Log(Enc(consumerMap2))

}

func TestCreateMapHooks(t *testing.T) {
	// one map with hooks
	var consumerMap2 = map[string]any{
		"AppID": 33,
	}
	Err(t, txHooks().Table(`tbl_consumer`).Create(&consumerMap2).Error)
	// create with user_id
	t.Log(Enc(consumerMap2))

	// map list with hooks
	var consumerMapList = []map[string]any{}
	for i := 0; i < 100; i++ {
		consumerMapList = append(consumerMapList, map[string]any{
			"AppID": (i + 1) * 5,
		})
	}
	Err(t, txHooks().Table(`tbl_consumer`).Create(&consumerMapList).Error)
	// created with user_id and after create hidden user_id
	t.Log(Enc(consumerMapList))

}
