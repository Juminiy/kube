package kstmt

type DistinctFromExpr struct {
	Start        Expr
	Not          *Empty
	DistinctFrom *Empty
	End          Expr
}
