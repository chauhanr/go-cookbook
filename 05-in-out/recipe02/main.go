package main

import (
	"io"
	"os"
	"fmt"
)

func main(){
	io.WriteString(os.Stdout, "This is string to standard output.\n")
	io.WriteString(os.Stdout, "This is string to standard error.\n")

	buf := []byte{0xAF, 0xFF, 0xFE}

	for i:=0; i<20; i++{
		if _, e:= os.Stdout.Write(buf); e != nil{
			panic(e)
		}
	}
	fmt.Fprintln(os.Stdout, "\n")
}
