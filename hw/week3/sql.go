package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Give a CSV file with the header:
	//
	// id,first,last,dept,salary
	//
	// create a sql database an issue a CREATE TABLE command to store the data.

	// Due to the strange issues we had with the IDE and the sql package, we'll
	// use the short declarations
	db, err := sql.Open("sqlite3", "week3.db")
	if err != nil {
		log.Fatal(err)
	}

	// Change the table name and the column names/types
	sql := "CREATE TABLE IF NOT EXISTS t1(a TEXT, b TEXT)"
	_, err = db.Exec(sql)
	// What always goes here?

	// Change the table name and the column names and count
	sql = "INSERT INTO t1(a,b) VALUES(?,?)"
	stmt, err := db.Prepare(sql)
	// ?? - Get used to it, every

	// Change the values passed to Exec
	_, err = stmt.Exec("a", "b")
	// single

	// Change the values passed to Exec
	_, err = stmt.Exec("c", "d")
	// time

	// Does this return an error?
	db.Close()
}
