package multi_tenants

import "gorm.io/gorm"

// ParseSchema
// experimental function
func (cfg *Config) ParseSchema(tx *gorm.DB) {
	stmt := tx.Statement
	if cfg.UseTableParseSchema && stmt.Schema == nil && len(stmt.Table) != 0 {
		if zeroV, ok := cfg.cacheStore.Load(stmt.Table); ok {
			if err := stmt.Parse(zeroV); err != nil {
				tx.Logger.Info(stmt.Context, "use table parse schema error: %s", err.Error())
			}
		}
	}
}
