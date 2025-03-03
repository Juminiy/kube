package gorm_api

import (
	"testing"
)

func TestPluck(t *testing.T) {
	var appID []uint
	Err(t, txPure().Model(&Consumer{}).
		Limit(10).
		Pluck("app_id", &appID).
		Error)
	t.Log(Enc(appID))
}

func TestCount(t *testing.T) {
	var cntConsumer int64
	Err(t, txPure().Model(&Consumer{}).
		Limit(10).
		Count(&cntConsumer).Error)
	t.Log(Enc(cntConsumer))
}

func TestMapColumn(t *testing.T) {
	var consumerMap map[string]any
	Err(t, txPure().Model(&Consumer{}).
		Limit(10).
		MapColumns(map[string]string{
			"id":        "$consumer_id",
			"app_id":    "$scope_id",
			"tenant_id": "$rented_id",
			"user_id":   "$consumer_account_id",
			"region":    "$country_in",
		}).
		First(&consumerMap).Error)
	t.Log(Enc(consumerMap))

	var consumerAlias struct {
		ID       uint
		AppID    uint
		TenantID uint
		UserID   uint
		Region   string
	}
	Err(t, txPure().Model(&Consumer{}).
		Limit(10).
		MapColumns(map[string]string{
			"id":        "consumer_id",
			"app_id":    "scope_id",
			"tenant_id": "rented_id",
			"user_id":   "consumer_account_id",
			"region":    "country_in",
		}).
		First(&consumerAlias).
		Error)
	t.Log(Enc(consumerAlias))
}
