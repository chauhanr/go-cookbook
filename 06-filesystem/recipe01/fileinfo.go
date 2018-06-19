package main

import (
	"os"
	"fmt"
)

func main(){

	f, err := os.Open("testfile.txt")
	if err != nil{
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil{
		panic(err)
	}

	fmt.Printf("File Name: %s\n", fi.Name())
	fmt.Printf("Is file with name %s a directory?: %v\n", fi.Name(), fi.IsDir())
	fmt.Printf("Size : %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())
}



