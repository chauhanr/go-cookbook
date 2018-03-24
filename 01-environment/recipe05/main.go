package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main(){
	ex, err := os.Executable()
	if err != nil{
		panic(err)
	}
	// path to the executable
	fmt.Println(ex)

	// resolving the directory of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path : "+exPath)

	// using eval symlinks
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil{
		panic(err)
	}
	fmt.Println("Symlink evaluated : "+realPath)
}