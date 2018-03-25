package main

import (
	"strconv"
	"fmt"
)

const binary = "00001"
const hex = "2f"
const intString = "12"
const floatString = "13.9"


func main(){
	res, err := strconv.Atoi(intString)  // converts array of character to integer therefore Atoi
	if err != nil{
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", res)

	// parsing hexdecimals
	res64, err := strconv.ParseInt(hex,16,32)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Parsed hexa decimal number: %d\n",res64)

	// parsing binary
	bin, err := strconv.ParseInt(binary,2,32)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Parsed binary value to number: %d\n", bin)

	// parsing floating point
	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Parsed binary value to number: %.5f\n", resFloat)

}

