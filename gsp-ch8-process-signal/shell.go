package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

const(
	VERSION = 0.2
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scanner.Scan(){
		line := scanner.Text()
		words := strings.Split(line," ")
		command := words[0]

		switch command {
		case "version":
			fmt.Println(VERSION)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(line)
		}
		fmt.Print("$ ")
	}

}
