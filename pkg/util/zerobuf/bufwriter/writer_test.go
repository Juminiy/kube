package bufwriter

import "testing"

func TestWriter_String(t *testing.T) {
	sqlBuf := New().
		Words("SELECT").WordsSep(", ", "name", "desc", "id").Line().
		Words("FROM", "`tbl_makabaka_like`").Line().
		Words("WHERE", "deleted_at", "IS", "NULL").Line().
		Words("AND", "sugar_id", "=").Worda(1).Line().
		Words("AND", "quarter_id", "IN").Byte('(').WordsaSep(", ", 1, 2, 3).Byte(')').Line()

	t.Log(sqlBuf.String())
}
