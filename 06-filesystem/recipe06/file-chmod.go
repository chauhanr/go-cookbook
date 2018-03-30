package main

import (
	"os"
	"fmt"
)

func main(){
	f, err := os.Create("testfile.txt")
	if err != nil{
		panic(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil{
		panic(err)
	}
	fmt.Printf("File permission %v\n", fi.Mode())

	err = f.Chmod(0777)
	if err != nil{
		panic(err)
	}
	fmt.Printf("File permission is %v\n", fi.Mode())
}
