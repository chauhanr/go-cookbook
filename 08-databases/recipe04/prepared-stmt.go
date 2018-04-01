package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const truncate = "truncate table post;"
const insert = "insert into post (ID, TITLE, CONTENT) values ($1, $2, $3)"

var testTable = []struct{
	ID int
	Title string
	Content string
}{
	{1,"Title One", "Content one"},
	{2, "Title Two", "Content two"},
	{3, "Title Three", "Content three"},
}

func main(){
	db := createConnection()
	_, err := db.Exec(truncate)
	if err != nil{
		panic(err)
	}
	stm, err := db.Prepare(insert)
	if err != nil{
		panic(err)
	}
	inserted := int64(0)
	for _, val := range testTable{
		fmt.Printf("Inserting records ID: %d\n", val.ID)
		r, err := stm.Exec(val.ID, val.Title, val.Content)
		if err != nil{
			fmt.Printf("Could not insert the record with id: %d \n", val.ID)
		}
		if affected, err := r.RowsAffected(); err == nil{
			inserted = inserted+affected
		}
	}
	fmt.Printf("Result: inserted %d rows. \n", inserted)
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