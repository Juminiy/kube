package kstmt

type OrderByClause AtLeastOne[OrderingTerm]

type OrderingTerm struct {
	Expr       Expr
	Collate    *string
	Order      *Order
	NullsFirst *Empty
	NullsLast  *Empty
}
