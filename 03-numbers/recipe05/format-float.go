package main

import "fmt"

var integer int64 = 32500
var floatNum float64 = 22000.456

func main(){
		// most common way to print integers
		fmt.Printf("%d \n", integer)
		// if we need to show the sign of the integer
		fmt.Printf("%+d \n", integer)

		// to print other bases x - base 16, o-8, b-2, d-10
		fmt.Printf("%X \n", integer)
		fmt.Printf("%#X \n", integer)

		// padding with leading zeros
		fmt.Printf("%010d \n", integer)
		// padding with leading spaces - left
		fmt.Printf("% 10d \n", integer)
		// padding with spaces - right
		fmt.Printf("% -10d \n", integer)

        // floating point precision of 5
		fmt.Printf("%.5f \n", floatNum)
		// floating point represented as scientific notation
		fmt.Printf("%e \n", floatNum)

}
