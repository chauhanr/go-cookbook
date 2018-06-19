package main

import (
	"os/user"
	"log"
	"fmt"
)

func main(){
	usr, err := user.Current()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("The users home directory is: ", usr.HomeDir)
}
