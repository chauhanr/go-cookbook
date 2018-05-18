package main

import (
	"os"
	"fmt"
	"log"
)

var (
	fileInfo os.FileInfo
)

func main(){
	fileInfo, err := os.Stat("file.txt")
	if err != nil{
		log.Fatalf("Error in reading file %s: %s\n", "file.txt", err.Error())
	}
	fmt.Printf("File name is: %s\n", fileInfo.Name())
	fmt.Printf("File Size in bytes %d\n", fileInfo.Size())
	fmt.Printf("Permissions : %s\n", fileInfo.Mode())
	fmt.Printf("Last Modified: %s\n", fileInfo.ModTime())
	fmt.Printf("File is a directory: %v\n", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
}
