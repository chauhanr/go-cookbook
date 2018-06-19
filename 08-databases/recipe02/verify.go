package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
	"context"
)

func main(){
	connStr := "postgres://postgres:postgres@localhost:5432/examples?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Ping O.K")
	ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)
	err = db.PingContext(ctx)

	if err != nil{
		fmt.Println("Error: "+err.Error())
	}
	conn,err := db.Conn(context.Background())
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	err = conn.PingContext(context.Background())
	if err != nil{
		panic(err)
	}
	fmt.Println("connection Ping O.K")
}