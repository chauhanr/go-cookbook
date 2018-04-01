package main

import (
	"database/sql"
	"context"
	"time"
	"fmt"
	_ "github.com/lib/pq"
)

// making a very long running query
const sel = "select * from post cross join (select 1 from generate_series(1,1000000)) tbl"

func main() {
	db := createConnection()
	defer db.Close()

	ctx, canc := context.WithTimeout(context.Background(), 20*time.Microsecond)
	rows, err := db.QueryContext(ctx, sel)
	// cancel the query returning the function.
	canc()
	if err !=nil{
		fmt.Println(err)
		return
	}
	defer rows.Close()
	count := 0
	for rows.Next(){
		if rows.Err() != nil{
			fmt.Println(rows.Err())
			continue
		}
		count++
	}
	fmt.Printf("%d rows returned\n", count)
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

