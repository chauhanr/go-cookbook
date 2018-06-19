package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main(){
	for i:=1; i<=6; i++{
		f, err := os.Create(fmt.Sprintf("./test%d.txt", i))
		if err != nil{
			fmt.Println(err.Error())
		}
		f.Close()
	}


	m, err := filepath.Glob("./test[1-3].txt")
	if err != nil{
		panic(err)
	}

	for _,val := range m{
		fmt.Println(val)
	}
	// cleaup
	for i:=1; i<=6; i++{
		err := os.Remove(fmt.Sprintf("./test%d.txt", i))
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}
