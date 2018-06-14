package main

import (
	"log"
	"os"
	"io"
)

/**
	io.ReadFull() method will read the first 512 bytes of the file that you provide.
	if the file is not 512 bytes long we get an error. This is very useful to check the boot sector of the disk
	that mounts the OS. We can compare the content of the boot sector with a previous known good version to ensure
	that there is no modification to the boot sector.
*/

func main(){
	path := "/dev/sda1"
	log.Println("[+] Reading boot sector of "+path)
	file, err := os.Open(path)
	if err != nil{
		log.Fatal("Error: "+err.Error())
	}
	// read the file using ReadFull first 512 bytes are read.
	byteSlice := make([]byte, 512)
	numBytes, err := io.ReadFull(file, byteSlice)
	if err != nil{
		log.Fatal("Unable to read 512 bytes of the file. "+ err.Error())
	}
	log.Printf("Bytes read: %d\n\n", numBytes)
	log.Printf("Data as String: \"%s\"\n",byteSlice)
	log.Printf("Data as Decimal: %d\n",byteSlice)
	log.Printf("Data as Hexa: %x\n",byteSlice)
}
