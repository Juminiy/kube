package kstmt

type ReturningClause struct {
	Expr AtLeastOne[ReturningExpr]
}

type ReturningExpr struct {
	Star *Empty
	Expr *ColumnAliasExpr
}

type ColumnAliasExpr struct {
	Expr        Expr
	As          *Empty
	ColumnAlias *string
}
