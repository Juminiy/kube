package kstmt

type BetweenAndExpr struct {
	Expr  Expr
	Not   *Empty
	LExpr Expr
	RExpr Expr
}
