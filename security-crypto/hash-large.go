package main

import (
	"fmt"
	"os"
		"log"
	"crypto/md5"
	"io"
	)

/*
	The hashing of large file cases an issue when trying to load all the contents of the file to memory.
	As hashing functions are block level cyphers we can us the io.Copy method that will copy the file in chunks.
    if we need to control the number of bytes we bring into the memory we can use the io.CopyBuffer() methods where
    you can define the block size. By default the block size is 32 jkb.
*/

func printUsageL(){
	fmt.Printf("Usage: %s <filepath>\n", os.Args[0])
	fmt.Printf("Example: %s document.txt", os.Args[0])
}

func checkArgsL() string{
	if len(os.Args) < 2{
		printUsageL()
		os.Exit(1)
	}
	return os.Args[1]
}

func main(){
	filePath := checkArgsL()
	file, err := os.Open(filePath)
	if err != nil{
		log.Fatal(err.Error())
	}
	defer file.Close()
	hasher := md5.New()
	// now coy the files in blocks and hand it over to the hasher
	_, err = io.Copy(hasher, file)
	if err != nil{
		log.Fatal(err)
	}
	checksum := hasher.Sum(nil)
	fmt.Printf("Md5 checksum %x\n", checksum)
}
