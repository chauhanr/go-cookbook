package main

import (
	"net/http"
	"fmt"
)

func main(){
	header := http.Header{}

	header.Set("Auth-X", "abcde2120fg")
	header.Add("Auth-X", "defgm34gh45")
	fmt.Println(header)

	// retrive values
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	resFirst := header.Get("Auth-X")
	fmt.Printf("The first value for Auth-X is :%s\n", resFirst)

	// replace all values of AuthX with new value
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// remove header
	header.Del("Auth-X")
	fmt.Println(header)

}