package main

import (
	"os"
	"log"
)

func main(){
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connection string: %s\n", connStr)
}