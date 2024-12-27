package kstmt

import "io"

type SelectStmt FactoredSelectStmt

func (b SelectStmt) Build(w io.Writer) {
}

type FactoredSelectStmt struct {
	SimpleSelectStmt
	Compound []CompoundSelect
}

type CompoundSelectStmt struct {
	SimpleSelectStmt
	Compound AtLeastOne[CompoundSelect]
}

type SimpleSelectStmt struct {
	With    *WithCommonTableExpr
	Select  SelectCore
	OrderBy *OrderByClause
	Limit   *LimitClause
}

type SelectCore struct {
	Expr   *SelectExpr
	Values *Values
}

type SelectExpr struct {
	Distinct *Empty
	All      *Empty
	Column   ResultColumnList
	From     FromExpr
	Where    *Expr
	GroupBy  *ExprList
	Having   *Expr
	Window   *AtLeastOne[WindowExpr]
}

type CompoundOperator struct {
	Union     *Empty
	UnionAll  *Empty
	Intersect *Empty
	Except    *Empty
}

type CompoundSelect struct {
	Operator CompoundOperator
	Select   SelectCore
}

type TableOrSubquery struct {
	Schema    *string
	TableExpr *struct {
		Name  string
		Index IndexExpr
	}
	TableFunc *TableFuncExpr
	Select    *SelectStmt
	As        *AsAliasExpr
	Recursive *struct {
		R    *AtLeastOne[TableOrSubquery]
		Join *JoinClause
	}
}
