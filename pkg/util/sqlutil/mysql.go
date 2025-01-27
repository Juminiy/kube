package sqlutil

import (
	tiparser "github.com/pingcap/tidb/pkg/parser"
	tiast "github.com/pingcap/tidb/pkg/parser/ast"
	timodel "github.com/pingcap/tidb/pkg/parser/model"
	_ "github.com/pingcap/tidb/pkg/parser/test_driver"
	"github.com/samber/lo"
)

func MySQLAst(sql string) (tiast.StmtNode, error) {
	stmt, _, err := tiparser.New().ParseSQL(sql)
	if err != nil {
		return nil, err
	}
	return stmt[0], nil
}

type ColumnVisitor struct {
	column []string
}

func (v *ColumnVisitor) Enter(in tiast.Node) (tiast.Node, bool) {
	if column, ok := in.(*tiast.ColumnName); ok {
		v.column = append(v.column, column.Name.String())
	}
	return in, false
}

func (v *ColumnVisitor) Leave(in tiast.Node) (tiast.Node, bool) {
	return in, true
}

func (v *ColumnVisitor) Column() []string {
	return v.column
}

type SelectColumnVisitor struct {
	column []ColumnRep
}

type ColumnRep struct {
	Name   *tiast.ColumnName
	AsName timodel.CIStr
}

func (v *SelectColumnVisitor) Enter(in tiast.Node) (tiast.Node, bool) {
	if fieldList, ok := in.(*tiast.FieldList); ok {
		v.column = lo.Map(fieldList.Fields, func(item *tiast.SelectField, i int) ColumnRep {
			var name tiast.ColumnName
			if item.WildCard != nil {
				name = tiast.ColumnName{
					Schema: item.WildCard.Schema,
					Table:  item.WildCard.Table,
				}
			} else {
				switch columnExpr := item.Expr.(type) {
				case *tiast.ColumnNameExpr:
					name = *columnExpr.Name
				}
			}
			return ColumnRep{
				Name:   &name,
				AsName: item.AsName,
			}
		})
	}
	return in, false
}

func (v *SelectColumnVisitor) Leave(in tiast.Node) (tiast.Node, bool) {
	return in, true
}

func (v *SelectColumnVisitor) Column() []ColumnRep {
	return v.column
}

type InsertValueVisitor struct {
	value [][]any
}

func (v *InsertValueVisitor) Enter(in tiast.Node) (tiast.Node, bool) {
	if value, ok := in.(*tiast.ValuesExpr); ok {
		_ = value
	}
	return in, false
}

func (v *InsertValueVisitor) Leave(in tiast.Node) (tiast.Node, bool) {
	return in, true
}

func (v *InsertValueVisitor) Values() [][]any {
	return v.value
}
