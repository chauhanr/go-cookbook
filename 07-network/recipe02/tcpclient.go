package main

import (
	"net/http"
	"net"
	"io"
	"bufio"
	"time"
	"fmt"
	"context"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, rq *http.Request){
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server{
	return http.Server{
		Addr: addr,
		Handler: StringServer("Hello Gopher!"),
	}
}

var addr = "localhost:7070"

func main(){
	s := createServer(addr)
	go s.ListenAndServe()

	conn, err := net.Dial("tcp", addr)
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost:localhost:7070\r\n\r\n")
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(time.Second))
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(ctx)
}