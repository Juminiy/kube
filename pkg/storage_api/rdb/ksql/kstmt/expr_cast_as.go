package kstmt

type CastAsExpr struct {
	Expr     Expr
	TypeName TypeName
}

type TypeName struct {
	Name AtLeastOne[string]
	One  *SignedNumber
	Pair *Pair[SignedNumber]
}

type SignedNumber struct {
	Positive *Empty
	Negative *Empty
	Number   NumericLiteral
}
