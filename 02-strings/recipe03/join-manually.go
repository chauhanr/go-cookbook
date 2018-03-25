package main

import (
	"strings"
	"fmt"
)

const selectBase = "SELECT * FROM user WHERE %s"

var refStringSlice = []string{
	" FRIST_NAME = 'Jack' ",
	" INSURANCE_NO = 33445442",
	" EFF_FROM = SYSDATE",
}

type JoinFunc func(p string) string

func main(){
	joinFunc := func(p string) string{
		if strings.Contains(p, "INSURANCE"){
			return "OR"
		}
		return "AND"
	}

	result := JoinWithFunc(refStringSlice, joinFunc)
	fmt.Println(selectBase+result)
}

func JoinWithFunc(refStringSlice []string, joinFunc JoinFunc) string{
	concatenate := refStringSlice[0]
	for _, val := range refStringSlice[1:]{
		concatenate = concatenate+ joinFunc(val) + val
	}
	return concatenate
}

