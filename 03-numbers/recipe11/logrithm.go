package main

import (
	"math"
	"fmt"
)

func main(){
	// natural log
	ln := math.Log(math.E)
	fmt.Printf("Ln(E) = %.4f\n", ln)

	//log base 10
	log10 := math.Log10(100)
	fmt.Printf("Log10(100) = %.4f\n", log10)

	// log base 2
	log2 := math.Log2(2)
	fmt.Printf("Log2(2) = %.4f\n", log2)

	// log 6 to base 3
	log3_6 := Log(3,9)
	fmt.Printf("Log3(9) = %.4f\n", log3_6)
}

// for base > 1 and  x greater than 0
func Log(base, x float64) float64{
	return math.Log(x) / math.Log(base)
}


