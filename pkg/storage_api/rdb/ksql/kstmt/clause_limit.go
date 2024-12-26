package kstmt

type LimitClause struct {
	Expr   Expr
	Offset *Expr
	Comma  *Expr
}
