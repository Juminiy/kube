package kstmt

type DeleteStmt struct {
	With      *WithCommonTableExpr
	Table     QualifiedTable
	Where     *Expr
	Returning *ReturningClause
	Limit     *DeleteStmtLimited
}

type DeleteStmtLimited struct {
	Delete  DeleteStmt
	OrderBy *OrderByClause
	Limit   *LimitClause
}
