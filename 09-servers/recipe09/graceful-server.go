package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"os/signal"
	"context"
	"time"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w,"hello world")
	})

	srv := &http.Server{Addr:":8080",  Handler: mux}

	go func(){
		if err := srv.ListenAndServe(); err != nil{
			log.Printf("Server error: %s\n", err.Error())
		}
	}()
	log.Printf("Server listening on: "+srv.Addr)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan // blocks the processing
	log.Printf("Shutting down the server gracefully ...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	srv.Shutdown(ctx)
	<- ctx.Done()
	cancel()
	log.Println("Server gracefully shutdown")
}
