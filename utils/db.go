package utils

import (
	"database/sql"
	"github.com/go-lib-utils-master/time"
	_ "github.com/go-sql-driver/mysql"
	time2 "time"
)

var (
	Db *sql.DB
	err error
)

func init()  {
	Db,err = sql.Open("mysql","root:lll2002.11.22@tcp(127.0.0.1:3306)/cinema")
	if err != nil {
		panic(err)
	}
}

func TestTime(timeData string) int {
	past,_:=time.ParseDateTime(timeData)
	now := time2.Now()

	if now.After(past) {
		return 0
	}
	return 1
}