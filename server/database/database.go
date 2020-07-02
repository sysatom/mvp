package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var DB *sqlx.DB

func Connect() {
	DB = sqlx.MustConnect("mysql", os.Getenv("MYSQL_URL"))
}
