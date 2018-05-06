package main

import (
	"crypto/sha256"
	"fmt"
	"time"
	"flag"
	"os"
)

func main(){

	length := flag.Int("l",8, "length of the password if not given 8 characters are assumed.")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1{
		fmt.Println("Usage gen-passwd [Options] seed")
		os.Exit(-1)
	}
	seed := args[0]
    sha := sha256.New()
	_, err := sha.Write([]byte(time.Now().String()+seed))
	if err != nil{
		fmt.Printf("Err while encoding sha %s", err.Error())
	}
	pwd := fmt.Sprintf("%x", sha.Sum(nil))
	if *length <= len(pwd){
		fmt.Printf("%s\n", pwd[len(pwd)-*length:])
	}else{
		fmt.Printf("Error lenght of the passwd cannot be more than %s", len(pwd))
	}


}