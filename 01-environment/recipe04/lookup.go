package main

import (
	"os"
	"log"
	"fmt"
)

func main(){
	args := os.Args

	if len(args) == 1{
		log.Printf("Usage : lookup env1 env2 env3....")
	}

	if len(args)> 1 {
		for _, key := range args[1:] {
			connStr, exists := os.LookupEnv(key)
			if !exists {
				log.Printf("Then env variable %s is not set.\n", key)
			}

			fmt.Printf("Env variable %s == %s\n",key, connStr)
		}
	}

}