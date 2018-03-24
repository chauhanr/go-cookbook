package main

import (
	"fmt"
	"strings"
	"flag"
	"log"
	"os"
)

type ArrayValue []string

func (a *ArrayValue) String() string{
	return fmt.Sprintf("%v",*a)
}

func (a *ArrayValue) Set(s string) error{
	*a = strings.Split(s,",")
	return nil
}

func main(){

	/*Trying to retrieve flags*/
	retry := flag.Int("retry", -1, "Defines max retry count")

	// read the flag useing the XXXVar function
	var logPrefix string
	flag.StringVar(&logPrefix,"prefix", "", "Logger prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array to iterate through")

	/* Flag parse function will parse the input sent to the program based on the rules defined earlier */
	flag.Parse()

	// this is logic that will use the flags.
	logger := log.New(os.Stderr, logPrefix, log.Ldate)
	retryCount := 0
	for retryCount < *retry{
		logger.Printf("Retrying connection")
		logger.Printf("Sending array %v\n",arr)
		retryCount++
	}
}

// possible calls from the command line will be   ./util  -retry 2 -prefix=example -array=1,2