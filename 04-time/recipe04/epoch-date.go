package main

import (
	"time"
	"fmt"
)

func main(){
	t := time.Unix(0,0)
	fmt.Println(t)
	// getting the epoch
	epoch := t.Unix()
	fmt.Println(epoch)

	apocNow := time.Now().Unix()
	fmt.Printf("Epoch time in seconds: %d \n", apocNow)

	apocNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano seconds: %d\n", apocNano)
}