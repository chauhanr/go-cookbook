package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main(){
	arguments := os.Args
	if len(arguments) ==1 {
		fmt.Println("Please provide a valid argument!")
		os.Exit(1)
	}
	filename := arguments[1]

	fileinfo, err := os.Lstat(filename)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	if fileinfo.Mode() & os.ModeSymlink != 0 {
		fmt.Println(filename, "is a symbolic link")
		realpath, err := filepath.EvalSymlinks(filename)
		if err == nil{
			fmt.Println("Path: ", realpath)
		}
	}else{
		fmt.Println(filename, "is not a symbolic link")
	}
}