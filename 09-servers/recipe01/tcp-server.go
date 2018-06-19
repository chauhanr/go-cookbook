package main

import (
	"net"
	"fmt"
	"bufio"
	"io"
)

func main(){
	l, err := net.Listen("tcp", ":8080")
	if err != nil{
		panic(err)
	}
	for {
		fmt.Println("Wating for client...")
		conn, err := l.Accept()
		if err != nil{
			panic(err)
		}
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil{
			panic(err)
		}
		_, err = io.WriteString(conn, "Received: "+string(msg))
		if err != nil{
			panic(err)
		}
		conn.Close()
	}
}
