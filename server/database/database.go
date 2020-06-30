package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() {
	DB = sqlx.MustConnect("mysql", "root:123456@(127.0.0.1:3318)/mvp?charset=utf8mb4&parseTime=true")
}
