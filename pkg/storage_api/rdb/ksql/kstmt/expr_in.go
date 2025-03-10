package kstmt

type InExpr struct {
	Expr       Expr
	Not        *Empty
	SelectStmt *SelectStmtExpr
	Schema     *string
	TableExpr  TableExpr
}

type TableExpr struct {
	Table     *string
	TableFunc *TableFuncExpr
}

type SelectStmtExpr struct {
	SelectStmt *SelectStmt
	Expr       ExprList
}
