package main

import (
	"fmt"
	"os"
	"bytes"
	"bufio"
	"io/ioutil"
)

func main(){
	fmt.Println("### Read as a reader###")
	f, err := os.Open("temp/file.txt")
	if err != nil{
		panic(err)
	}
	defer f.Close()
	// read from the file reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan(){
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### Read File ###")
	fContent, err := ioutil.ReadFile("temp/file.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(fContent))
}
