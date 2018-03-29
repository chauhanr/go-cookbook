package main

import (
	"os"
	"fmt"
)

func main(){
	for {
		data := make([]byte,8)
		n, err := os.Stdin.Read(data)

		if err == nil && n>0{
			process(data)
		}else{
			break
		}
	}
}

func process(data []byte){
	fmt.Printf("Received data %s\n", string(data))
}