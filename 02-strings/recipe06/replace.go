package main

import (
	"strings"
	"fmt"
)

const refString = "Mary had a little lamb"

func main(){
	// this will use replacer packge
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out := replacer.Replace(refString)
	fmt.Println(out)

	// now using the replace method to revert a change
	out = strings.Replace(out, "wolf", "lamb", -1)
	fmt.Println(out)
}
