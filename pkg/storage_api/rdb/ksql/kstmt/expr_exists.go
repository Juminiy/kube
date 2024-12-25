package kstmt

type ExistsExpr struct {
	Exists     *NotExistsExpr
	SelectStmt SelectStmt
}

type NotExistsExpr struct {
	Not *Empty
}
