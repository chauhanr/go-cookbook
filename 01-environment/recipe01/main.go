package main

import (
	"log"
	"runtime"
)

const info=`
Application %s starting. 
The binary was build by go version: %s
`

func main(){
	log.Printf(info, "recipe03", runtime.Version())
}