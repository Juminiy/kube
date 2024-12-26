package kstmt

type Column struct {
	Name  string
	Table *Table
}

type Table struct {
	Name   string
	Schema *string
}

type TableFunc string

type TableFuncExpr struct {
	Func TableFunc
	Expr ExprList
}

type ColumnList AtLeastOne[string]

type IndexedColumn struct {
	Column  ColumnExpr
	Collate *string
	Order   *Order
}

type ColumnExpr struct {
	Column *string
	Expr   *Expr
}

type Order struct {
	Asc  *Empty
	Desc *Empty
}

type QualifiedTable struct {
	Table Table
	As    *AsAliasExpr
	Index *IndexExpr
}

type IndexExpr struct {
	IndexedBy  *string
	NotIndexed *Empty
}

type ResultColumn struct {
	Expr         *ReturningExpr
	TableDotStar *string
}

type ResultColumnList AtLeastOne[ResultColumn]
