package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const sel = "SELECT * FROM post;"
const truncate = "TRUNCATE TABLE post;"
const insert = "INSERT INTO post (ID, TITLE, CONTENT) VALUES (1, 'TITLE 1', 'CONTENT 1'), (2, 'TITLE 2', 'CONTENT 2') "

func main(){
	db := createConnection()
	_, err := db.Exec(truncate)
	if err != nil{
		panic(err)
	}
	fmt.Println("Table truncated.")
	r, err := db.Exec(insert)
	if err != nil{
		panic(err)
	}
	affected, err := r.RowsAffected()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Inserted row count: %d \n", affected)
	rs, err := db.Query(sel)
	if err != nil{
		panic(err)
	}
	count :=0
	for rs.Next(){
		count++
	}
	fmt.Printf("Total of %d rows were retrieved \n", count)
}

func createConnection() *sql.DB{
	connStr := "postgres://postgres:postgres@localhost:5432/examples?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		panic(err)
	}
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	return db
}
