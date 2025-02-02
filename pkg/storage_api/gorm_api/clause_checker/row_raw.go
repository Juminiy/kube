package clause_checker

import "gorm.io/gorm"

func (cfg *Config) RowRawClause(tx *gorm.DB) {
	if _, ok := tx.Get(SkipRawOrRow); ok || tx.Error != nil {
		return
	}

	if cfg.AllowWriteClauseToRawOrRow {
		cfg.WhereClause(tx)
		if where, ok := WhereClause(tx); ok {
			_, _ = tx.Statement.WriteString(" WHERE ")
			where.Build(tx.Statement)
		}

		if limit, ok := LimitClause(tx); ok {
			_ = tx.Statement.WriteByte(' ')
			limit.Build(tx.Statement)
		}
	}
}
