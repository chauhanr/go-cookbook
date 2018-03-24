package main

import (
	"os"
	"fmt"
)

func main(){
	 /* first we print all the arguments passed to the program.*/
	 args := os.Args
	 fmt.Println(args)

	 // the first argument is always the name of the binary.
	 fmt.Printf("The binary name is %s\n", args[0])

	 if len(args) >= 2{
		 // we iterate through the rest of the parameters.
		 for i, arg := range args[1:] {
			 fmt.Printf("agrument %d is %s\n",i+1, arg)
		 }
	 }else{
	 	fmt.Printf("There were no arguments passed to this program")
	 }
}