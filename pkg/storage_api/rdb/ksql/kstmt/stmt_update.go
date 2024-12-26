package kstmt

type UpdateStmt struct {
	With      *WithCommonTableExpr
	Opt       *WriteFailOpt
	Table     QualifiedTable
	Set       UpdateSetExpr
	From      *FromExpr
	Where     *Expr
	Returning *ReturningClause
}

type UpdateSetExpr AtLeastOne[SetExpr]

type SetExpr struct {
	Column  *string
	Columns *ColumnList
	Expr    Expr
}

type FromExpr struct {
	TableOrSubquery *AtLeastOne[TableOrSubquery]
	Join            *JoinClause
}
