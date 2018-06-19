package main

import (
	"fmt"
	"crypto/md5"
	"os"
	"io"
)

var content = "This is the content to check"

func main(){
	checksum := MD5(content)
	checksum2 := FileMD5("content.dat")

	fmt.Printf("Checksum of content %s \nChecksum of file %s\n", checksum, checksum2)
	if checksum == checksum2 {
		fmt.Printf("Two checksums are same!")
	}
}

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func FileMD5(path string) string{
	h := md5.New()
	f, err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	_,err = io.Copy(h,f)
	if err != nil{
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
