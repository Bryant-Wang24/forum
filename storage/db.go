package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("mysql", "root:485969746wqs@(localhost:3306)/forum")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to mysql success")
}
