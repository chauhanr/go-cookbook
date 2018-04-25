package main

import (
	"os"
	"fmt"
	"flag"
	"errors"
	"path/filepath"
)

func main(){
	list := flag.Bool("l", false, "list details")

	flag.Parse()
	err := errors.New("")

	path := ""
	arguments := flag.Args()
	if len(arguments) == 0{
		path = "."
	}else{
		path = arguments[0]
	}
	switch {
		case *list:
			err = filepath.Walk(path, walkList)
		default:
			err = filepath.Walk(path, walkFuncNoOption)
	}

	fmt.Println("")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func walkFuncNoOption(path string, info os.FileInfo, err error) error{
	_, err = os.Stat(path)
	if err != nil{
		return err
	}
	fmt.Printf("%s ",path)
	return nil
}

func walkList(path string, info os.FileInfo, err error) error{
	_, err = os.Stat(path)
	if err != nil{
		return err
	}
	fmt.Println(path)
	return nil
}