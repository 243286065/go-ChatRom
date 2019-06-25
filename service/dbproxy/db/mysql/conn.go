package mysql

import (
	"database/sql"
	"fmt"
	cfg "go-ChatRom/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dbErr error

func init() {
	db, dbErr = sql.Open("mysql", cfg.MySQLSource)

	if dbErr != nil {
		fmt.Println(dbErr.Error())
		os.Exit(1)
	}

	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func DBConn() *sql.DB {
	return db
}
