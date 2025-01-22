package main

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/chdb-io/chdb-go/chdb"
	_ "github.com/chdb-io/chdb-go/chdb/driver"
	"log"
)

const DefaultFormat = "json"

func main() {
	ses, err := chdb.NewSession("chdb.io")
	util.Must(err)
	//defer ses.Close()

	/*	crtRes, err := ses.Query(`
		CREATE TABLE IF NOT EXISTS default.my_first_table
		(
		    user_id UInt32,
		    message String,
		    timestamp DateTime,
		    metric Float32
		)
		ENGINE = MergeTree
		PRIMARY KEY (user_id, timestamp)
		`, DefaultFormat)
			if err != nil {
				log.Printf("create table error: %s", err.Error())
				return
			}
			log.Println(crtRes.String())*/

	insRes, err := ses.Query(`
INSERT INTO default.my_first_table (user_id, message, timestamp, metric) VALUES
    (101, 'Hello, ClickHouse!',                                 now(),       -1.0    ),
    (102, 'Insert a lot of rows per batch',                     yesterday(), 1.41421 ),
    (102, 'Sort your data based on your commonly-used queries', today(),     2.718   ),
    (101, 'Granules are the smallest chunks of data read',      now() + 5,   3.14159 )
`, DefaultFormat)
	if err != nil {
		log.Printf("insert table error: %s", err.Error())
		return
	}
	log.Println(insRes.String())

	qryRes, err := ses.Query(`
SELECT *
FROM default.my_first_table
WHERE user_id >= 10 AND user_id <= 110
ORDER BY timestamp
`, DefaultFormat)
	if err != nil {
		log.Printf("do query error: %s", err.Error())
		return
	}
	log.Println(qryRes.String())
}
