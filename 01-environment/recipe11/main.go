package main

import (
	"os"
	"fmt"
	"time"
	"log"
	"io"
	"os/signal"
	"syscall"
)

var writer *os.File
func main(){
	/**first open a log will that the server will write to*/
	var err error

	writer, err = os.OpenFile(fmt.Sprintf("test_%d.log", time.Now().Unix()), os.O_RDWR | os.O_CREATE, os.ModePerm)
	if err != nil{
		panic(err)
	}

	// now open a channel and allwo the server to do it jobs.
	closeChan := make(chan bool)

	go func(){
	  for {
		  time.Sleep(time.Second)
		  select {
		  case <-closeChan:
			  log.Printf("Go routine closing \n")
			  return
		  default:
			  log.Printf("Writing a log to file.\n	")
			  io.WriteString(writer, fmt.Sprintf("Logging access at : %s\n", time.Now().String()))
		  }
	  }
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGINT)
	// here we block the server till we get a signal for termination
	<- sigChan
	// all statements after the signal are clean up statements.
	close(closeChan)
	releaseAllResources()
	fmt.Printf("Shutting down the server gracefully")
}

func releaseAllResources(){
	io.WriteString(writer, "Application releasing all resources it holds\n")
	writer.Close()
}
