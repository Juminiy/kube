package kstmt

type CaseWhenExpr struct {
	Expr     *Expr
	WhenThen AtLeastOne[WhenThenExpr]
	Else     *ElseExpr
}

type WhenThenExpr struct {
	When Expr
	Then Expr
}

type ElseExpr struct {
	Expr Expr
}
