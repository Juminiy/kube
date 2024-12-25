package kstmt

type JoinClause struct {
	TableOrSubquery TableOrSubquery
	Joins           []Join
}

type TableOrSubquery struct {
}

type Join struct {
	Operator   JoinOperator
	Dest       TableOrSubquery
	Constraint JoinConstraint
}

type JoinOperator struct {
	Comma *Empty
	Join  *JoinOp
}

type JoinOp struct {
	Cross *Empty
	Op    *JoinOp0
}

type JoinOp0 struct {
	Natural *Empty
	Op      *JoinOp1
}

type JoinOp1 struct {
	Inner *Empty
	Op    *JoinOp2
}

type JoinOp2 struct {
	Scope *LRFOp
	Outer *Empty
}

type LRFOp struct {
	Left  *Empty
	Right *Empty
	Full  *Empty
}

type JoinConstraint struct {
	On    *Expr
	Using *ColumnList
}
