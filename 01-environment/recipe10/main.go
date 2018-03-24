package main

import (
	"os/exec"
	"bufio"
	"fmt"
	"io"
	"time"
)

func main(){
	cmd := []string{"go", "run", "sample.go"}
	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	stdIn, _ := proc.StdinPipe()
	defer stdIn.Close()

	stdOut,_ := proc.StdoutPipe()
	defer stdOut.Close()

	go func(){
		s := bufio.NewScanner(stdOut)
		for s.Scan(){
			fmt.Printf("%s\n", s.Text())
		}
	}()

	proc.Start()
	// now we are writing to the command line
	fmt.Printf("Writing input!\n")
	io.WriteString(stdIn, "Hello\n")
	io.WriteString(stdIn, "Golang\n")
	io.WriteString(stdIn, "is awesome\n")

	time.Sleep(time.Second*4)
	proc.Process.Kill()
}