package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func printUsage(){
	fmt.Printf("Usage: %s <filepath>\n", os.Args[0])
	fmt.Printf("Example: %s document.txt", os.Args[0])
}

func checkArgs() string{
	if len(os.Args) < 2{
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

func main(){
   filepath := checkArgs()
   data, err := ioutil.ReadFile(filepath)
   if err != nil {
   	   log.Fatal("Error reading file "+filepath)
   }

   fmt.Printf("Md5: %x\n", md5.Sum(data))
   fmt.Printf("sha1: %x\n", sha1.Sum(data))
   fmt.Printf("sha256: %x\n", sha256.Sum256(data))
   fmt.Printf("sha512: %x\n", sha512.Sum512(data))
}
