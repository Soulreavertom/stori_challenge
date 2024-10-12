package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	dburl := os.Getenv("DBURL")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	confdbstr := user + ":" + pass + "@tcp(" + dburl + ":" + dbport + ")/" + dbname
	//DB, err = sql.Open("mysql", "user:password@tcp(url:port)/dbname")
	DB, err = sql.Open("mysql", confdbstr)

	if err != nil {
		//panic(fmt.Sprintf("Error opening database: %v", err))
		fmt.Printf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		//panic(fmt.Sprintf("Error connecting to the database: %v", err))
		fmt.Printf("Error connecting to the database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

}
