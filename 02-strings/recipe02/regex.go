package main

import (
	"fmt"
	"regexp"
)

const refStringRegex = "Mary*had,a%little/lamb"

func main(){

	words := regexp.MustCompile("[*,%_/]{1}").Split(refStringRegex, -1)
	for idx, word := range words{
		fmt.Printf("Word %d is: %s\n",idx, word)
	}

}