package mysqlwrapper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB

func InitPool() error {
	var err error
	dsn := "xilan:xilan123@tcp(localhost:3306)/wtestgpt"
	pool, err = sql.Open("mysql", dsn)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	return err
}

func ReleasePool() {
	pool.Close()
}

type TableUserRecord struct {
	userid          string
	total_use_count int
	version         int
}

func GetUserByUserId(userid string) (*TableUserRecord, error) {
	rows, err := pool.Query("SELECT * FROM user WHERE userid = $1", userid)
	if err != nil {
		return nil, err
	}
	var rec TableUserRecord
	rows.Next()
	err = rows.Scan(&rec.userid, &rec.total_use_count, &rec.version)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", rec)
	return &rec, nil
}
