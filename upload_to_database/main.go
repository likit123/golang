/*
Example upload file .TEXT to Database
Text file size 2.7GB , Upload to database use about 3 min
*/
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {

	now_begin := time.Now().Format("2006-01-02 15:04:05")

	var err error

	db, err = sqlx.Open("mysql", "_USER_:_PASSWORD_@tcp(_IPADDRESS_:3306)/_DATABASE_")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("truncate table TABLE")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("update var_log set begin=? where name='bi_report_transit_clr003_all'", now_begin)
	if err != nil {
		log.Fatal(err)
	}

	file, error := os.Open("BI_REPORT_TRANSIT_CLR003_ALL_20220801.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()

	_, err = db.Exec(`LOAD DATA INFILE '/_PATH_/BI_REPORT_TRANSIT_CLR003_ALL_20220801.txt' INTO TABLE bi_report_transit_clr003_all FIELDS TERMINATED BY '^' LINES TERMINATED BY '\n' IGNORE 1 ROWS`)

	if err != nil {
		log.Fatal(err)
	}

	now_end := time.Now().Format("2006-01-02 15:04:05")

	_, err = db.Exec("update var_log set end=? where name='bi_report_transit_clr003_all'", now_end)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(now_end)

}
