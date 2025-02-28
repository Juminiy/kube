package multi_tenants

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (cfg *Config) tenantClause(tx *gorm.DB, forQuery bool) {
	if tInfo := cfg.TenantInfo(tx); tInfo != nil {
		tInfo.AddClause(tx)
		if !forQuery {
			tx.Omit(tInfo.Field.DBName)
		}
	}
}

func (t *Tenant) AddClause(tx *gorm.DB) {
	tx.Statement.AddClause(t)
}

func (t *Tenant) Name() string { return "Tenant" }

func (t *Tenant) Build(_ clause.Builder) {}

func (t *Tenant) MergeClause(_ *clause.Clause) {}

// referred from: gorm.SoftDeleteQueryClause
func (t *Tenant) ModifyStatement(stmt *gorm.Statement) {
	if c, ok := stmt.Clauses["WHERE"]; ok {
		if where, ok := c.Expression.(clause.Where); ok && len(where.Exprs) >= 1 {
			for _, expr := range where.Exprs {
				if orCond, ok := expr.(clause.OrConditions); ok && len(orCond.Exprs) == 1 {
					where.Exprs = []clause.Expression{clause.And(where.Exprs...)}
					c.Expression = where
					stmt.Clauses["WHERE"] = c
					break
				}
			}
		}
	}

	stmt.AddClause(clause.Where{Exprs: []clause.Expression{
		t.Field.Clause(),
	}})
}
