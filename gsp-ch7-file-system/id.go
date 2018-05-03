package main

import (
	"os"
	"fmt"
	"os/user"
)

func main(){
	arguments := os.Args

	if len(arguments) == 1{
		uid := os.Getuid()
		fmt.Printf("Guid for user is: %d", uid)
		return
	}

	username := arguments[1]
	u, err := user.Lookup(username)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(u.Uid)
}

