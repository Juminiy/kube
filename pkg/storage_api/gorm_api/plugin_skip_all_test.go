package gorm_api

import (
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/plugin_register"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"sync"
	"testing"
)

var once sync.Once

func skipTx() *gorm.DB {
	once.Do(func() {
		util.Must(plugin_register.OneError(
			txPure().Callback().Create().Remove("multi_tenants:before_create"),
			txPure().Callback().Create().Remove("multi_tenants:after_create"),
			txPure().Callback().Create().Remove("multi_tenants:before_query"),
			txPure().Callback().Create().Remove("multi_tenants:after_query"),
			txPure().Callback().Create().Remove("multi_tenants:before_update"),
			txPure().Callback().Create().Remove("multi_tenants:after_update"),
			txPure().Callback().Create().Remove("multi_tenants:before_delete"),
			txPure().Callback().Create().Remove("multi_tenants:after_delete"),
			txPure().Callback().Create().Remove("multi_tenants:before_row"),
			txPure().Callback().Create().Remove("multi_tenants:after_row"),
			txPure().Callback().Create().Remove("multi_tenants:before_raw"),
			txPure().Callback().Create().Remove("multi_tenants:after_raw"),
		))
	})
	return txPure()
}

func TestSkipPlugin(t *testing.T) {
	consumerMap := map[string]any{
		"AppID": 10,
	}
	Err(t, skipTx().
		Model(&Consumer{}).
		Create(&consumerMap).Error)
	t.Log(Enc(consumerMap))

	// panic
	/*consumerMap2 := map[string]any{
		"AppID": 20,
	}
	Err(t, skipTx().
		Table(`tbl_consumer`).
		Create(&consumerMap2).Error)
	t.Log(Enc(consumerMap2))*/
}
