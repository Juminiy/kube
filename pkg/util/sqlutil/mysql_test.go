package sqlutil

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	tiast "github.com/pingcap/tidb/pkg/parser/ast"
	"testing"
)

func fullTestFunc(t *testing.T, sql string) {
	stmtAst, err := MySQLAst(sql)
	util.Must(err)

	// which kind of statement
	t.Logf("SQL Type: %s", tiast.GetStmtLabel(stmtAst))

	// get all exist columns
	colVis := &ColumnVisitor{
		column: make([]string, 0, util.MagicSliceCap),
	}
	stmtAst.Accept(colVis)
	t.Logf("All Column: %v", colVis.Column())

	switch tiast.GetStmtLabel(stmtAst) {
	case "Select":
		// get select columns
		selVis := &SelectColumnVisitor{}
		stmtAst.Accept(selVis)
		t.Logf("Select Column: %v", safe_json.Pretty(selVis.Column()))

	case "Insert":
		// get insert values
		insVis := &InsertValueVisitor{}
		stmtAst.Accept(insVis)
		t.Logf("Insert Values: %v", safe_json.Pretty(insVis.Values()))
	}

}

func TestMySQLParserSelect(t *testing.T) {
	fullTestFunc(t, "SELECT `id`,`name`,`desc`,`code`,`price` FROM `tbl_product` WHERE ( id = 1 AND name LIKE \"\" OR id = 1 OR id = 2 AND NOT id = 3 AND NOT id = 4 AND `tbl_product`.`tenant_id` = 114514) AND `tbl_product`.`deleted_at` IS NULL ORDER BY id desc,id asc,id DESC,id ASC LIMIT 10")
}

func TestMySQLParserInsert(t *testing.T) {
	fullTestFunc(t, `
INSERT INTO tbl_sys_realm (c, c, b)
VALUES 	(0721, 'MyGO', '2025-01-15 13:25:48'),
	   	(1919810, 'K-ON', '2023-05-31 19:45:33'),
		(114514, 'CCB', '2018-01-26 22:15:21')
`)
}
