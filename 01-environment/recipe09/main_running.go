package main

import (
	"runtime"
	"os/exec"
	"fmt"
)

func main(){
	/** first check the type of os */
	var cmd string
	if runtime.GOOS == "windows"{
		cmd = "timeout"
	}else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "2")
	proc.Start()
	fmt.Printf("Process State for running process is : %v\n", proc.ProcessState)

	fmt.Printf("PID for the process : %d \n", proc.ProcessState.Pid)
}


