package kstmt

import "io"

type CommonTableExpr struct {
	Table  string
	Column []string

	Materialized *struct {
		Not *Empty
	}

	SelectStmt SelectStmt
}

func (b CommonTableExpr) Build(w io.Writer) {

}
