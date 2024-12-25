package kstmt

type LikeOrGRMExpr struct {
	Expr Expr
	Not  *Empty
	Like *LikeExpr
	GRM  *GRMExpr
}

type LikeExpr struct {
	Expr   Expr
	Escape *EscapeExpr
}

type EscapeExpr struct {
	Expr Expr
}

type GRMExpr struct {
	Glob   *Empty
	Regexp *Empty
	Match  *Empty
	Expr   Expr
}
