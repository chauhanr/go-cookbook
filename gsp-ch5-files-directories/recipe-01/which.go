package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var DEBUG bool = false

func debug(msg string){
	if DEBUG {
		fmt.Println(msg)
	}
}

func main(){
	opA := flag.Bool("a", false, "this lists all the search - all")
	opS := flag.Bool("s", false, "s")
	Debug := flag.Bool("v", false, "verbose logging")

	flag.Parse()
	flags := flag.Args()
	if len(flags) == 0{
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}
	DEBUG = *Debug

	for _, file := range flags{
		foundIt := whichFunc(file, opS, opA)

		if foundIt == false{
			fmt.Printf("Could not find executable %s in path\n", file)
		}
	}

}

func whichFunc(file string, opS, opA *bool) bool{
	foundIt := false

	path := os.Getenv("PATH")
	debug("Path : "+path)
	pathSplice := strings.Split(path, ":")

	debug(fmt.Sprintf("Path Splices are :%v", pathSplice))

	for _, dir := range pathSplice{
		fullPath := dir +"/" +file
		fileinfo, err := os.Stat(fullPath)
		if err == nil{
			mode := fileinfo.Mode()
			debug(fmt.Sprintf("File: %s found and mode is %s", fileinfo.Name(), mode.String()))
			if mode.IsRegular() {
				if mode&0111 != 0{
					foundIt = true
					if *opS ==true{
						os.Exit(0)
					}
					if *opA == true{
						fmt.Printf("%s : %s\n",fullPath, fileinfo.Mode().String())
					}else {
						fmt.Println(fullPath)
						return foundIt
					}
				}
			}
		}else{
			debug(fmt.Sprintf("Could not find file %s", fullPath))
		}
	}
	return foundIt
}