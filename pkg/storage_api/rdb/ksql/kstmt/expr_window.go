package kstmt

type WindowExpr struct {
	Name string
	As   WindowDefn
}

type WindowDefn struct {
	BaseName    *string
	PartitionBy *ExprList
	OrderBy     *OrderByClause
	Frame       *FrameSpec
}

type FrameSpec struct {
	RRG     FrameSpecRRG
	BUEC    FrameSpecBUEC
	Exclude *FrameSpecExclude
}

type FrameSpecRRG struct {
	Range  *Empty
	Rows   *Empty
	Groups *Empty
}

type FrameSpecBUEC struct {
	BetweenAnd         *FrameSpecBetweenAnd
	UnboundedPreceding *Empty
	ExprPreceding      *Expr
	CurrentRow         *Empty
}

type FrameSpecBetweenAnd struct {
	LExpr struct {
		UnboundedPreceding *Empty
		ExprPreceding      *Expr
		CurrentRow         *Empty
		ExprFollowing      *Expr
	}
	RExpr struct {
		ExprPreceding      *Expr
		CurrentRow         *Empty
		ExprFollowing      *Expr
		UnboundedFollowing *Empty
	}
}

type FrameSpecExclude struct {
	NoOthers   *Empty
	CurrentRow *Empty
	Group      *Empty
	Ties       *Empty
}
