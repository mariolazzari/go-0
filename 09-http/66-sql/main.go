package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// select
	rows, err := db.Query("Select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}

	// insert
	stmt, err := db.Prepare("insert into users(name) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Mario")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last ID:", lastId)
}
