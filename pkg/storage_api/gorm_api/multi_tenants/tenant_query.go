package multi_tenants

import (
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	gormschema "gorm.io/gorm/schema"
	"slices"
)

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil || _SkipQueryCallback.Ok(tx) {
		return
	}
	if GetSessionConfig(cfg, tx).BeforeQueryOmitField {
		beforeQueryOmit(tx)
	}

	cfg.tenantClause(tx, true)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := GetSessionConfig(cfg, tx)

	if sCfg.AfterFindMapCallHooks {
		afterFindMapCallHook(tx)
	}

	if !sCfg.AfterQueryShowTenant {
		if tInfo := cfg.TenantInfo(tx); tInfo != nil {
			_Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.Name: nil, // FieldName
			})
		}
	}
}

var _SkipQueryCallback = Cfg{
	key: "skip_query_callback",
}

func beforeQueryOmit(tx *gorm.DB) {
	// replaced by gorm tag `->:false` is an AfterQuery Set Fields To Zero
	// QueryOmit is Omit query column
	if sch := tx.Statement.Schema; sch != nil {
		slices.All(sch.Fields)(func(_ int, field *gormschema.Field) bool {
			if !field.Readable {
				tx.Omit(field.DBName)
			}
			return true
		})
	}
}

func afterFindMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.AfterFind {
		/*setUpDestMapStmtModel(db, sch)*/
		CallHooks(db, func(v any, tx *gorm.DB) bool {
			if afterFindI, ok := v.(callbacks.AfterFindInterface); ok {
				_ = db.AddError(afterFindI.AfterFind(tx))
				return true
			}
			return false
		})
		scanModelToDestMap(db)
	}
}
