package main

import (
	"os"
	"io"
	"crypto/md5"
	"fmt"
	"bufio"
)

var data = []struct{
	name string
	cont string
	perm os.FileMode
}{
	{"test1.file" , "Hello Golang is great", 0666},
	{"test2.file" , "Hello Golang is great", 0666},
	{"test3.file" , "Not matching\n Golang is great", 0666},
}

func main(){
	files := []*os.File{}
	for _, fData := range data{
		f, err := os.Create(fData.name)
		if err != nil{
			panic(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, fData.cont)
		if err != nil{
			panic(err)
		}
		files = append(files, f)
	}

	checksums := []string{}
	for _, f := range files {
		f.Seek(0,0 )
		sum, err := getMD5SumString(f)
		if err != nil{
			panic(err)
		}
		checksums = append(checksums,sum)
	}

	fmt.Println("#### Comparing by checksum ###")
	compareChecksum(checksums[0], checksums[1])
	compareChecksum(checksums[0], checksums[2])

	fmt.Println("### Comparing by line ###")
	files[0].Seek(0,0)
	files[2].Seek(0,0)
	compareByLine(files[0], files[2])

	// clean up
	for _, val := range data{
		os.Remove(val.name)
	}

}

func compareChecksum(s1, s2 string){
 	match := "match"
 	if s1 != s2 {
 		match = "does not match"
	}
	fmt.Printf("Sum %s and Sum %s %s\n", s1, s2, match)
}


func getMD5SumString(f *os.File) (string, error){
	fileSum := md5.New()
	_,err := io.Copy(fileSum,f)
	if err != nil{
		return "", err
	}
	return fmt.Sprintf("%X", fileSum.Sum(nil)), nil
}

func compareByLine(f1, f2 *os.File){
	sc1 := bufio.NewScanner(f1)
	sc2 := bufio.NewScanner(f2)

	for {
		sc1Bool := sc1.Scan()
		sc2Bool := sc2.Scan()
		if !sc1Bool && !sc2Bool{
			break
		}
		compareLines(sc1.Text(), sc2.Text())
	}
}

func compareLines(line1, line2 string){
	sign := "o"
	if line1 != line2 {
		sign = "x"
	}
	fmt.Printf("%s | %s | %s \n", sign, line1, line2)
}