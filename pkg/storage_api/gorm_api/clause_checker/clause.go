package clause_checker

import "gorm.io/gorm"

// gorm support clause
const (
	Delete     = "DELETE"
	From       = "FROM"
	GroupBy    = "GROUP BY"
	Insert     = "INSERT"
	Limit      = "LIMIT"
	Locking    = "FOR"
	OnConflict = "ON CONFLICT"
	OrderBy    = "ORDER BY"
	Returning  = "RETURNING"
	Select     = "SELECT"
	Set        = "SET"
	Update     = "UPDATE"
	Values     = "VALUES"
	Where      = "WHERE"
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
