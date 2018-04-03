package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

const call = "select * from format_name($1, $2, $3)"
const callMySQL = "CALL sampleproc(?)"

type Result struct{
	Name string
	Category string
}

func main(){
	db := createConnection()
	defer db.Close()
	r := Result{}

	if err := db.QueryRow(call, "Ritesh", "Chauhan", 38).Scan(&r.Name); err != nil{
		panic(err)
	}
	fmt.Printf("Result is: %+v\n", r)
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