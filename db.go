package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func connectDB() {
	fmt.Println("Database connected")
	db, err = sql.Open("mysql","root:Abcd@1234@/server_go")
}
