package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	// scanner
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan(){
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}
}
