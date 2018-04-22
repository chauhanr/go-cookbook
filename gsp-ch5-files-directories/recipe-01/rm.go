package main

import (
	"os"
	"fmt"
)

func main(){
	arguments := os.Args
	if len(arguments) == 1{
		fmt.Println("Please provide an argument")
		os.Exit(1)
	}
	file := arguments[1]
	err := os.Remove(file)
	if err != nil{
		fmt.Printf("%s\n", err.Error())
		return
	}else{
		fmt.Println("successfully deleted file ", file)
	}
}
