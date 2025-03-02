package gorm_api

import (
	"gorm.io/gorm"
	"testing"
)

func TestCallbacksBeforeDelete(t *testing.T) {
	Err(t, txMixed().Delete(&Consumer{
		Model: gorm.Model{ID: 2},
		AppID: 20,
	}).Error)
}
