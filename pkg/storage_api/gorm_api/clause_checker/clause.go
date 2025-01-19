package clause_checker

import "gorm.io/gorm"

// gorm support clause
const (
	Where      = "WHERE"
	Returning  = "RETURNING"
	OnConflict = "ON CONFLICT"
	From       = "FROM"
	Set        = "SET"
	Select     = "SELECT"
	Limit      = "LIMIT"
	OrderBy    = "ORDER BY"
	GroupBy    = "GROUP BY"
)

func (cfg *Config) Clause(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	for _, fn := range []func(tx *gorm.DB){
		cfg.WhereClause,
		cfg.OrderByClause,
		cfg.LimitClause,
		cfg.GroupByClause,
	} {
		fn(tx)
	}
}
