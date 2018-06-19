package main

import (
	time2 "time"
	"context"
	"os/exec"
	"bufio"
	"fmt"
)

func main(){
	cmd := "ping"
	timeout := 2 * time2.Second

    ctx, _ := context.WithTimeout(context.TODO(), timeout)
    proc := exec.CommandContext(ctx, cmd,"www.google.com")

    stdOut,_ := proc.StdoutPipe()
    defer stdOut.Close()

    proc.Start()

    s := bufio.NewScanner(stdOut)
    for s.Scan(){
    	fmt.Println(s.Text())
	}
}