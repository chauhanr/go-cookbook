package main

import (
	"fmt"
	"math"
)

var val float64 = 3.55554444

func main(){
	intVal := int(val)
	fmt.Printf("Bad rounding by casting to int :%d\n", intVal)
	fRound := Round(val)
	fmt.Printf("Rouding by custom function: %v\n", fRound)
}

func Round(val float64) float64{
	t := math.Trunc(val)
	if math.Abs(val-t) >= 0.5{
		return t + math.Copysign(1,val)
	}
	return t
}