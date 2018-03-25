package main

import (
	"regexp"
	"fmt"
)

const refString = ` [{ \"email\": \"email@example.com\" ,
						\"phone\": 555234556}, 
					  { \"email\": \"ritesh@example.com\" ,
						\"phone\": 8800439536}]`

func main(){
	emailRegex := regexp.MustCompile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}")
	first := emailRegex.FindString(refString)
	fmt.Printf("Fisrt: %s\n", first)

	findAll := emailRegex.FindAllString(refString,-1)
	fmt.Println("All Emails:")
	for _, val := range findAll{
		fmt.Println(val)
	}
}
