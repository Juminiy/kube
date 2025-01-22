package main

import (
	"database/sql"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/pkg/errors"
	"log"
)

func insertAndQuery() {
	db, err := sql.Open("chdb", "")
	if err != nil {
		log.Fatal(err)
	}
	cli := Click{db: db}

	var lastID, rowAff int64
	_ = cli.Exec(
		"INSERT INTO my_first_table (user_id, message, timestamp, metric) VALUES "+
			"(101, 'Hello, ClickHouse!',                                 now(),       -1.0    ), "+
			"(102, 'Insert a lot of rows per batch',                     yesterday(), 1.41421 ), "+
			"(102, 'Sort your data based on your commonly-used queries', today(),     2.718   ), "+
			"(101, 'Granules are the smallest chunks of data read',      now() + 5,   3.14159 )",
		&lastID, &rowAff)

	var uRecords []struct {
		UserID    uint
		Message   string
		Timestamp sql.NullTime
		Metric    float64
	}
	err = cli.Query(
		"SELECT * FROM my_first_table"+
			"WHERE id >= 100 AND id <= 110", &uRecords)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return
	}
	log.Println(safe_json.Pretty(uRecords))
}

type Click struct {
	db *sql.DB
}

func (c *Click) Exec(str string, lastID, rowAff *int64) error {
	res, err := c.db.Exec(str)
	if err != nil {
		return err
	}
	*lastID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	*rowAff, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (c *Click) Query(str string, res any) error {
	rows, err := c.db.Query(str)
	if err != nil {
		return errors.Wrap(err, "select fail")
	}
	cols, err := rows.Columns()
	if err != nil {
		return errors.Wrap(err, "get result columns fail")
	}
	log.Printf("%v", cols)
	defer util.SilentCloseIO("rows", rows)
	for rows.Next() {
		err := rows.Scan(res)
		if err != nil {
			return errors.Wrap(err, "scan fail")
		}
		log.Printf("%v", res)
	}
	return nil
}
