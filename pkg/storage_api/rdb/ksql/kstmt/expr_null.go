package kstmt

type NullExpr struct {
	Expr         Expr
	IsNull       *Empty
	NotNull      *Empty // NOTNULL
	NotSpaceNull *Empty // NOT NULL
}
