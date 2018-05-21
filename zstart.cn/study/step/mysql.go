// mysql
package step

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func StudyMysql() {
	testDB()
}

func testDB() {
	db, err := sql.Open("mysql", "root:hazel110@tcp(127.0.0.1:3306)/sakila?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select * from actor")
	var id, first, last, time []byte
	for rows.Next() {
		err = rows.Scan(&id, &first, &last, &time)
		checkErr(err)

		fmt.Println(string(id), string(first), string(last))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
