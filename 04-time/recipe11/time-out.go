package main

import (
	"time"
	"fmt"
)

func main(){
	to := time.After(3*time.Second)
	list := make([]string,0)

	done := make(chan bool, 1)

	fmt.Println("Starting the record insertion process")
	go func(){
		defer fmt.Println("Exiting go routine")
		for {
			select {
			case <-to:
				fmt.Println("The time is up")
				done <- true
			default:
				list = append(list, time.Now().String())
			}
	    }
	}()
	<-done
	fmt.Printf("Managed to insert %d items \n", len(list))
}
