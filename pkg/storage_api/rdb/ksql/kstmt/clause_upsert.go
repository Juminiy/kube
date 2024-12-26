package kstmt

type UpsertClause AtLeastOne[UpsertExpr]

type UpsertExpr struct {
	Nothing *Empty
}

type ConflictTarget struct {
}
