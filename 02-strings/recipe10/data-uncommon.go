package main

import (
	"os"
	"encoding/csv"
	"fmt"
)

func main(){
	file, err := os.Open("data-uncommon.csv")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord=3
	reader.Comment='#'
	reader.Comma=';'

	for {
		record, e := reader.Read()
		// record is nothing but an array of strings.
		if e != nil{
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
