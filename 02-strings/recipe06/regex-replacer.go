package main

import (
	"regexp"
	"fmt"
)

const refStrignReg = "Mary had a little lamb"

func main(){
	regex := regexp.MustCompile("l[a-z]+") // will match any word that starts with l
	out := regex.ReplaceAllString(refStrignReg,"replacement")
	fmt.Println(out)
}
