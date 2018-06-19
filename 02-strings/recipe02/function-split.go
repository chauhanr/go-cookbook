package main

import (
	"strings"
	"fmt"
)

const refStringFunc = "Mary*had,a%little/lamb"

func main(){

	// define the split function that determines if we have a delimiter
	splitFunc := func(r rune) bool{
		return strings.ContainsRune("%*_,/", r)
	}

	words := strings.FieldsFunc(refStringFunc, splitFunc)
	for idx, word := range words{
		fmt.Printf("Word %d is: %s\n",idx, word)
	}

}