package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan(){
		fmt.Println(sc.Text())
	}
}
