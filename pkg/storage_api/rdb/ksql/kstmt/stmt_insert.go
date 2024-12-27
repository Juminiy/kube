package kstmt

type InsertStmt struct {
	With   *WithCommonTableExpr
	Write  ReplaceOrInsertExpr
	Table  Table
	As     *AsAliasExpr
	Column ColumnList
	Data   DataFromStmt
}

type WithCommonTableExpr struct {
	Recursive *Empty
	TableExpr AtLeastOne[CommonTableExpr]
}

type ReplaceOrInsertExpr struct {
	Replace *Empty
	Insert  *InsertExpr
}

type InsertExpr struct {
	Opt *WriteFailOpt
}

type WriteFailOpt struct {
	Abort    *Empty
	Fail     *Empty
	Ignore   *Empty
	Replace  *Empty
	Rollback *Empty
}

type AsAliasExpr struct {
	As    *Empty
	Alias string
}

type DataFromStmt struct {
	Values    *Values
	Select    *SelectStmt
	Upsert    *UpsertClause
	Default   *Empty
	Returning *ReturningClause
}

type Values AtLeastOne[ExprList]
