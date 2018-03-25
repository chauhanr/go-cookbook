package main

import (
	"strings"
	"fmt"
)

const refStringDel = "Mary_had a little_lamb"

func main(){
	words := strings.Split(refStringDel, "_")
	for idx, word := range words{
		fmt.Printf("Word %d is: %s\n",idx, word)
	}
}

