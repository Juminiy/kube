package kstmt

type BinaryOperator string

type BinaryExpr struct {
	LExpr    Expr
	Operator BinaryOperator
	RExpr    Expr
}
