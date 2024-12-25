package kstmt

type FuncExpr struct {
	Name         string
	Args         FuncArgs
	FilterClause *FilterClause
	OverClause   *OverClause
}

type FuncArgs struct {
}

type FilterClause struct {
}

type OverClause struct {
}
