package main

import (
	"strings"
	"fmt"
)

const refString = "Mary had a little lamb"

/**
	This recipe helps you search if a word or phrase is present in the string.*/

func main() {

	searchParams := []string{"Mary", "Wolf", "lamb"}
	for _, lookFor := range searchParams {
		contain := strings.Contains(refString, lookFor)
		fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)
		hasPrefix := strings.HasPrefix(refString, lookFor)
		fmt.Printf("The \"%s\" has prefix \"%s\": %t \n", refString, lookFor, hasPrefix)
		hasSuffix :=  strings.HasSuffix(refString, lookFor)
		fmt.Printf("The \"%s\" has suffix \"%s\": %t \n\n", refString, lookFor, hasSuffix)
	}

}