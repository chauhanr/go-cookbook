package main

import (
	"os"
	"fmt"
	"flag"
	"path/filepath"
)

func main(){
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0{
		fmt.Println("Not enough arguments !")
		os.Exit(1)
	}
	Path := flags[0]
	err := filepath.Walk(Path, walk)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func walk(path string, info os.FileInfo, err error) error{
	fileInfo, err := os.Stat(path)
	if err != nil{
		return err
	}

	mode := fileInfo.Mode()
	if mode.IsDir() || mode.IsRegular(){
		fmt.Println(path)
	}

	return nil
}
