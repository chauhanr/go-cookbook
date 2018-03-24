package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main(){
	/*Create a channel that will receive os signals os/signals of size 1*/
	sChan := make(chan os.Signal,1)
	/*Next notify the channel for the following signals
	1. signal on closing of the terminal
	2. signal for calling ctl+c
	3. signal quit
	4. signal terminated process kill command*/
	signal.Notify(sChan, syscall.SIGHUP,
						syscall.SIGINT,
						syscall.SIGTERM,
						syscall.SIGQUIT)

    exitChan := make(chan int)

    go func(){
		signal := <- sChan
		switch signal{
		case syscall.SIGHUP:
			fmt.Println("The calling terminal was closed.")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process is terminated as it was interruptd by ctrl+c")
			exitChan <- 1
		case syscall.SIGTERM:
			fmt.Println("Kill SIGTERM has been initiated on the process")
			exitChan <- 1
		case syscall.SIGQUIT:
			fmt.Println("Kill SIGOUIT was executed for process")
			exitChan <- 1
		}
	}()

    code := <- exitChan
    os.Exit(code)
}
