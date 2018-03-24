package main

import (
	"os/exec"
	"bytes"
	"fmt"
)

func main(){
	prc := exec.Command("ls", "-lta")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	if prc.ProcessState.Success() {
		fmt.Println("Process ran successfully with output:")
		fmt.Println(out.String())
	}
}