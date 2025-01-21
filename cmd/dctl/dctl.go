package main

import (
	"fmt"
	"github.com/chdb-io/chdb-go/chdb"
	"os"
	//_ "github.com/chdb-io/chdb-go/chdb/driver"
)

func main() {
	if len(os.Args) == 3 {
		fmt.Println(chdb.Query(os.Args[1], os.Args[2]))
		return
	}

	//db, err := sql.Open("chdb", "")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//rows, err := db.Query(`select COUNT(*) from url('https://datasets.clickhouse.com/hits_compatible/athena_partitioned/hits_0.parquet')`)
	//if err != nil {
	//	log.Fatalf("select fail, err: %s", err)
	//}
	//cols, err := rows.Columns()
	//if err != nil {
	//	log.Fatalf("get result columns fail, err: %s", err)
	//}
	//log.Printf("result columns: %v", cols)
	//defer util.SilentCloseIO("rows", rows)
	//var count int
	//for rows.Next() {
	//	err := rows.Scan(&count)
	//	if err != nil {
	//		log.Fatalf("scan fail, err: %s", err)
	//	}
	//	log.Printf("count: %d", count)
	//}

}
