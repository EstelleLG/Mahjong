package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER        = "zimuwang"
	DB_PASSWORD    = "random"
	DB_NAME        = "postgres"
	DB_PRINT_LIMIT = 20
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("DB opened")

	fmt.Println("Querying")

	rows, err := db.Query("SELECT * FROM words")
	checkErr(err)
	count := 0
	for rows.Next() {
		var word string
		err = rows.Scan(&word)
		checkErr(err)
		count += 1
		if count < DB_PRINT_LIMIT {
			fmt.Printf("%s\n", word)
		}
	}

	fmt.Println("Inserting new record")

	var lastInsertWord string
	err = db.QueryRow("INSERT INTO words(word) VALUES($1) returning word;", "astaxie").Scan(&lastInsertWord)

	fmt.Printf("Last inserted word %s", lastInsertWord)
}
