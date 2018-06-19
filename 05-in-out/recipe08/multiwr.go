package main

import (
	"bytes"
	"os"
	"io"
	"fmt"
)

func main(){
	buf := bytes.NewBuffer([]byte{})
	f, err := os.OpenFile("sample.txt", os.O_CREATE | os.O_RDWR, os.ModePerm)
	if err != nil{
		panic(err)
	}
	wr := io.MultiWriter(buf, f)  // we have taken two writers and send them to the varidic function.
	_, err = io.WriteString(wr, "Hello, Go is awesome")
	if err != nil{
		panic(err)
	}
	fmt.Println("Content of biffer: "+buf.String())

}
