package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


const sel = "select * from post;"

func main(){
	db := createConnection()
	defer db.Close()

	rs, err := db.Query(sel)
	if err != nil{
		panic(err)
	}

	defer rs.Close()
	columns, err := rs.Columns()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Selected columns: %v\n", columns)

	colTypes, err := rs.Columns()
	if err != nil{
		panic(err)
	}
	for _, col := range colTypes{
		fmt.Println()
		fmt.Printf("%+v\n", col)
	}
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