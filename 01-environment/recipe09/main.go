package main

import (
	"runtime"
	"os/exec"
	"fmt"
	"time"
)

func main(){

	var cmd string
	if runtime.GOOS == "windows"{
		cmd = "timeout"
	}else {
		cmd = "sleep"
	}
    fmt.Printf("Running command %s\n",cmd )
	proc := exec.Command(cmd, "2")
	proc.Start()
	// we will wait fro the process to end
	proc.Wait()

	fmt.Printf("Process State for ended process : %v\n",proc.ProcessState)
	fmt.Printf("Process System time %dms\n", proc.ProcessState.SystemTime()/time.Microsecond )
	fmt.Printf("Did the process with pid : %d exit successfully %t\n",proc.ProcessState.Pid(), proc.ProcessState.Success())

}