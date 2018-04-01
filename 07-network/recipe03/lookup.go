package main

import (
	"net"
	"fmt"
)

func main(){
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil{
		panic(err)
	}
	fmt.Println("Resolving the address 127.0.0.1")
	for _,addr := range addrs{
		fmt.Println(addr)
	}
	fmt.Println("Resolving the domain localhost")
	ips, err := net.LookupIP("localhost")
	if err != nil{
		panic(err)
	}
	for _, ip := range ips{
		fmt.Println(ip.String())
	}
}
