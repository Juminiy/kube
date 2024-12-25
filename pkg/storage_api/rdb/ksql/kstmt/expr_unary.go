package kstmt

type UnaryOperator string

type UnaryExpr struct {
	Operator UnaryOperator
	Expr     Expr
}
