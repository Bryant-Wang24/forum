package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sqlx.DB
var gormDb *gorm.DB

func init() {
	var err error
	db, err = sqlx.Open("mysql", "root:485969746wqs@(localhost:3306)/forum?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	gormDb, err = gorm.Open(gorm_mysql.New(gorm_mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = gormDb.Exec(("select 1")).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to mysql success")
}
