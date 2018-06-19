package main

import (
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
)

func ListDirByWalk(path string){
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error{
		if wPath == path{
			return nil
		}
		if info.IsDir(){
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}
		return nil
	})
}

func ListByReadDir(path string){
	lst, err := ioutil.ReadDir(path)
	if err != nil{
		panic(err)
	}
	for _,val := range lst{
		if val.IsDir(){
			fmt.Printf("[%s]\n", val.Name())
		}else{
			fmt.Println(val.Name())
		}
	}
}

func main(){
	fmt.Println("List the content of the current dir")
	ListByReadDir("../")
	fmt.Println()
	fmt.Println("List by Walk")
	ListDirByWalk("../")
}