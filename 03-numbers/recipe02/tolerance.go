package main

import (
	"fmt"
	"math"
)

const da = 0.2999999999999999889776975374843459576368331909180
const db = 0.3

func main(){
	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	fmt.Printf("Strings %s = %s eqauls %v \n",daStr, dbStr, daStr == dbStr)
	fmt.Printf("Numbers equal: %v\n", db == da)

	// when comparing the floating points it is better to use a tolerance value
	fmt.Printf("Numbers equals with tolerance: %v\n", equals(da,db))
}

const TOLERANCE = 1e-8

func equals(a, b float64) bool{
	delta := math.Abs(a-b)
	if delta < TOLERANCE{
		return true
	}
	return false
}
