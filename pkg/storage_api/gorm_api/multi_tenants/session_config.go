package multi_tenants

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
)

type SessionConfig struct {
	DisableFieldDup          bool // effect on create and update
	ComplexFieldDup          bool // effect on create
	DeleteAllowTenantAll     bool // effect on delete, only tenant and soft_delete, no other where clause
	BeforeDeleteDoQuery      bool // effect on delete
	UpdateAllowTenantAll     bool // effect on update, only tenant and soft_delete, no other where clause
	UpdateOmitMapZeroElemKey bool // effect on update
	AfterCreateShowTenant    bool // effect on create
	AfterQueryShowTenant     bool // effect on query
	BeforeQueryOmitField     bool // effect on query
}

func GetSessionConfig(cfg *Config, tx *gorm.DB) SessionConfig {
	if v, ok := tx.Get(SessionCfg); ok {
		if vRecv, ok := v.(SessionConfig); ok {
			return vRecv
		} else if pRecv, ok := v.(*SessionConfig); ok && pRecv != nil {
			return *pRecv
		}
	}
	cfg.ParseSchema(tx)
	return *cfg.GlobalCfg
}

const SessionCfg = "session_config"

var _GlobalCfg = &SessionConfig{}

type Cfg struct {
	key string
	val any     // maybe unused, maybe nil
	do  util.Fn // maybe unused, maybe nil
}

func (c *Cfg) Set(tx *gorm.DB) *gorm.DB {
	return tx.Set(c.key, struct{}{})
}

func (c *Cfg) Ok(tx *gorm.DB) bool {
	_, ok := tx.Get(c.key)
	return ok
}

func (c *Cfg) Val(tx *gorm.DB) any {
	return c.val
}
