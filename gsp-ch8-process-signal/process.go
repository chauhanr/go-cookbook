package main

import (
	"os/exec"
	"fmt"
	"os"
	"syscall"
)

func main(){
	PS, err := exec.LookPath("ps")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(PS)
	command := []string{"ps", "-a", "-x"}
	env := os.Environ()

	err = syscall.Exec(PS, command, env)
	if err != nil{
		fmt.Printf("Error when executing command %s\n", PS)
	}
}

