package kstmt

type FuncExpr struct {
	Name         string
	Args         FuncArgs
	FilterClause *FilterClause
	OverClause   *OverClause
}

type FuncArgs struct {
	Distinct *Empty
	Expr     *ExprList
	OrderBy  *OrderByClause
	Star     *Empty
}

type FilterClause struct {
	Where Expr
}

type LiteralValue struct {
	Numeric          *NumericLiteral
	String           *StringLiteral
	Blob             *BlobLiteral
	Null             *Empty
	True             *Empty
	False            *Empty
	CurrentTime      *Empty
	CurrentDate      *Empty
	CurrentTimestamp *Empty
}

type NumericLiteral struct {
}

type StringLiteral struct {
}

type BlobLiteral struct {
}

type OverClause struct {
	WindowName *string
	Window     *WindowDefn
}

type RaiseFunc struct {
	Ignore       *Empty
	Rollback     *Empty
	Abort        *Empty
	Fail         *Empty
	ErrorMessage *string
}
