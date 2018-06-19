package main

import (
	"time"
	"fmt"
)

func main(){

	tTime := time.Date(2017,time.March,5,8,5,2,0,time.Local)
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))
	// time for hours and mins
	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))
	// predefined date format can be used as well.
	fmt.Printf("The time is: %s\n",tTime.Format(time.RFC1123))
	// zero padding
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))

}
