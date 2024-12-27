package kstmt

type UpsertClause AtLeastOne[UpsertExpr]

type UpsertExpr struct {
	Target  ConflictTarget
	Nothing *Empty
	Update  *struct {
		Set   *UpdateSetExpr
		Where *Expr
	}
}

type ConflictTarget struct {
	Column *AtLeastOne[IndexedColumn]
	Where  *Expr
}
