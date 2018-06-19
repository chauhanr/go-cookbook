package main

import (
	"net"
	"fmt"
)

func main(){

	interfaces, err := net.Interfaces()
	if err!= nil{
		panic(err)
	}
	for _, inter := range interfaces{
		addrs, err := inter.Addrs()
		if err!= nil{
			panic(err)
		}
		fmt.Println(inter.Name)
		for _, add := range addrs{
			if ip, ok := add.(*net.IPNet); ok{
				fmt.Printf("IP address: %v\n",ip)
			}
		}
		fmt.Printf("\n")
	}
}
