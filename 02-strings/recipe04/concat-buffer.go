package main

import (
	"bytes"
	"fmt"
)

func main(){
	strings := []string{"This ", "is ", "even ", "more ", "performant "}
	buf := bytes.Buffer{}

	for _, value := range strings{
		buf.WriteString(value)
	}
	fmt.Println(buf.String())
}

