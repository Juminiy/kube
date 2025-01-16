package multi_tenants

import "gorm.io/gorm"

type SessionConfig struct {
	DisableFieldDup          bool // effect on create and update
	ComplexFieldDup          bool // effect on create
	DeleteAllowTenantAll     bool // effect on delete, only tenant and soft_delete, no other where clause
	QueryBeforeDelete        bool // effect on delete
	UpdateAllowTenantAll     bool // effect on update, only tenant and soft_delete, no other where clause
	UpdateOmitMapZeroElemKey bool // effect on update
	AfterCreateShowTenant    bool // effect on create
	AfterQueryShowTenant     bool // effect on query
}

func GetSessionConfig(cfg *Config, tx *gorm.DB) SessionConfig {
	if UseSession.Get(tx) {
		v := UseSession.Value(tx)
		if vRecv, ok := v.(SessionConfig); ok {
			return vRecv
		} else if pRecv, ok := v.(*SessionConfig); ok && pRecv != nil {
			return *pRecv
		}
	}
	return *cfg.SessionConfig
}

var UseSession = SingleConfig{
	Key: "session_config",
}

type SingleConfig struct {
	Key string
}

func (s SingleConfig) Set(tx *gorm.DB) *gorm.DB {
	return tx.Set(s.Key, struct{}{})
}

func (s SingleConfig) Get(tx *gorm.DB) bool {
	_, ok := tx.Get(s.Key)
	return ok
}

func (s SingleConfig) Value(tx *gorm.DB) any {
	v, _ := tx.Get(s.Key)
	return v
}
