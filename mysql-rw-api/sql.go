package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var writeDB *sql.DB
var readDB *sql.DB

func initDB() {
	var err error

	writeDSN := "root:rootpass@tcp(localhost:3307)/appdb?parseTime=true"
	writeDB, err = sql.Open("mysql", writeDSN)
	if err != nil {
		log.Fatal("write db open error:", err)
	}

	readDSN := "root:rootpass@tcp(localhost:3308)/appdb?parseTime=true"
	readDB, err = sql.Open("mysql", readDSN)
	if err != nil {
		log.Fatal("read db open error:", err)
	}

	writeDB.SetMaxOpenConns(20)
	writeDB.SetMaxIdleConns(10)

	readDB.SetMaxOpenConns(20)
	readDB.SetMaxIdleConns(10)

	// Verify connections
	if err = writeDB.Ping(); err != nil {
		log.Fatal("write db ping error:", err)
	}
	if err = readDB.Ping(); err != nil {
		log.Fatal("read db ping error:", err)
	}

	log.Println("Connected to primary and replica")
}
