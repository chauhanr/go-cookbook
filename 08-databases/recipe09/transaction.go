package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


const selOne = "select id, title, content from post where ID = $1;"
const insert = "insert into post (ID, TITLE, CONTENT) values (4, 'Transaction Title 1', 'Transaction Content 1');"

type Post struct{
	ID int
	Title string
	Content string
}

func main(){
	db := createConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil{
		panic(err)
	}
	_, err = tx.Exec(insert)
	if err != nil{
		panic(err)
	}
	p := Post{}
	if err := db.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil{
		fmt.Printf("Got Error for db query: %s\n", err.Error())
	}
	fmt.Println(p)
	// using the transaction
	if err = tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil{
		fmt.Printf("Got error for db query under transaction: %s\n", err.Error())
	}
	fmt.Println(p)
	tx.Rollback()
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
