package main

import "fmt"

func main(){
	strings := []string{"This ", "is ", "even ", "more ", "performant "}
	buf := make([]byte, 100)
	count := 0

	for _, val := range strings{
		count += copy(buf[count:], []byte(val))
	}
	fmt.Println(string(buf[:]))
}
