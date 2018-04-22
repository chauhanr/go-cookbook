package main

import (
	"os"
	"fmt"
	"path/filepath"
	"flag"
	"errors"
)

func main(){
	dirOnly := flag.Bool("d", false, "Traverse directories only")
	flag.Parse()

	arguments := flag.Args()
	if len(arguments) == 0{
		fmt.Println("Usage traverse [options] <directory>")
		os.Exit(1)
	}

	path := arguments[0]
	err := errors.New("")

	switch {
		case *dirOnly:
			err = filepath.Walk(path, walkDir)
		default:
			err = filepath.Walk(path, walkAll)
	}
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

/**
	this function is used as an input to the filepath.Walk function.
*/
func walkAll(path string, info os.FileInfo, err error) error{
	_, err = os.Stat(path)

	if err != nil{
		return err
	}
	fmt.Println(path)
	return nil
}

func walkDir(path string, info os.FileInfo, err error) error{
	file, err := os.Stat(path)
	if err != nil{
	return err
	}
	if file.IsDir(){
		fmt.Println(path)
	}
	return nil
}