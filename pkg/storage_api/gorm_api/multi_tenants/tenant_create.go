package multi_tenants

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeCreate(tx *gorm.DB) {
	tid, ok := cfg.tenantValid(tx)
	if !ok {
		return
	}

	stmt := tx.Statement
	refv := _Ind(stmt.ReflectValue)
	field, ok := util.MapElemOk(util.MapVK(refv.Tag1(cfg.TagKey)), cfg.TagTenantKey)
	if !ok {
		return
	}
	refv.SetField(map[string]any{
		field: tid,
	})
}

func (cfg *Config) AfterCreate(tx *gorm.DB) {

}
