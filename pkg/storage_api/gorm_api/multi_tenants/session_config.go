package multi_tenants

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
)

type SessionConfig struct {
	DisableFieldDup          bool // effect on create and update
	ComplexFieldDup          bool // effect on create
	DeleteAllowTenantAll     bool // effect on delete, if false: raise error when clause only have (tenant) and (soft_delete)
	BeforeDeleteDoQuery      bool // effect on delete, use with clause.Returning, when database not support Returning
	UpdateAllowTenantAll     bool // effect on update, if false: raise error when clause only have (tenant) and (soft_delete)
	UpdateOmitMapZeroElemKey bool // effect on update
	UpdateOmitMapUnknownKey  bool // effect on update
	UpdateMapSetPkToClause   bool // effect on update
	AfterCreateShowTenant    bool // effect on create
	BeforeQueryOmitField     bool // effect on query, use with tag `gorm:"->:false"`
	AfterQueryShowTenant     bool // effect on query

	// callbacks Hooks
	CreateMapCallHooks    bool // effect on create map
	UpdateMapCallHooks    bool // effect on update map
	AfterFindMapCallHooks bool // effect on find map
}

func GetSessionConfig(cfg *Config, tx *gorm.DB) SessionConfig {
	cfg.ParseSchema(tx)
	if v, ok := tx.Get(SessionCfg); ok {
		if vRecv, ok := v.(SessionConfig); ok {
			return vRecv
		} else if pRecv, ok := v.(*SessionConfig); ok && pRecv != nil {
			return *pRecv
		}
	}
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
