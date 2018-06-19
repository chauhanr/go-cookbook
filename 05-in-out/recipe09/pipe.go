package main

import (
	"os/exec"
	"log"
	"io"
	"os"
)

func main(){
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("echo", "Hello Go!\nThis is an example")
	cmd.Stdout = pWriter

	go func(){
		defer pReader.Close()
		if _, err := io.Copy(os.Stdout, pReader); err != nil{
			log.Fatal(err)
		}
	}()

	if err := cmd.Run(); err != nil{
		log.Fatal(err)
	}

}
