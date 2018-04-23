package main

import (
	"os"
	"fmt"
	"flag"
)

var (
	VERBOSE = false
)


func verboseLog(msg string){
	if VERBOSE{
		fmt.Println(msg)
	}
}

func main(){

	verbose := flag.Bool("v",false, "Verbose explains what the command is doing")
	flag.Parse()
	VERBOSE = *verbose

	arguments := flag.Args()
	if len(arguments) == 0{
		fmt.Println("Please provide an argument")
		os.Exit(1)
	}

	for _,file := range arguments{
		verboseLog(fmt.Sprintf("Deleting file %s", file))
		err := os.Remove(file)
		if err != nil{
			verboseLog(fmt.Sprintf("%s\n", err.Error()))
			return
		}else{
			verboseLog(fmt.Sprintf("successfully deleted file: %s\n", file))
		}
	}
}
