package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(username, dbname, ip, password string, port int) (msdb *sql.DB, err error) {
	msdb, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, ip, port, dbname))
	return
}
