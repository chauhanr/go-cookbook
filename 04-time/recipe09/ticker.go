package main

import (
	"os"
	"os/signal"
	"time"
	"fmt"
)

func main(){
	c := make(chan os.Signal,1)
	signal.Notify(c)

	ticker := time.NewTicker(time.Second)
	stop := make(chan bool)

	go func(){
		defer func() {
			stop <- true
		}()
		for {
			select {
				case <-ticker.C:
					fmt.Println("Tick")
				case <-stop:
					fmt.Printf("Goroutine closing\n")
					return
			}
		}
	}()

	<-c
	ticker.Stop()
	stop <- true
	<-stop
	fmt.Println("Application stoppeds")
}

