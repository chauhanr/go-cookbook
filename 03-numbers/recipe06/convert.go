package main

import (
	"strconv"
	"fmt"
)

const bin = "10111"
const hex = "1A"
const oct = "12"
const dec = "10"
const floatNum = 16.123557

func main(){
	// convert binary to hex
	v, _ := ConvertInt(bin, 2,16)
	fmt.Printf("Binary value %s converted to hex value : %s\n",bin,v)

	// convert hex to dec
	v, _ = ConvertInt(hex, 16,10)
	fmt.Printf("Hex value %s converted to Dec value : %s\n",hex,v)

	//convert oct to hex
	v, _ = ConvertInt(oct, 8,16)
	fmt.Printf("Oct value %s converted to Hex value : %s\n",oct,v)

	//convert from dec to oct
	v, _ = ConvertInt(dec, 10,8)
	fmt.Printf("Dec value %s converted to Oct value : %s\n",dec,v)
}

func ConvertInt(val string, fromBase int, toBase int) (string, error){
	i, err := strconv.ParseInt(val, fromBase, 64)
	if err != nil{
		return "", err
	}
	return strconv.FormatInt(i,toBase), nil
}
