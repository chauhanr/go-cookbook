package main

import (
	"os"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
)

func main(){
	if len(os.Args) != 4{
		fmt.Printf("usage : %s source destination BUFFERSIZE \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	source := os.Args[1]
	destination := os.Args[2]
	BUFFERSIZE, _ := strconv.ParseInt(os.Args[3], 10, 64)

	fmt.Printf("Copying file %s to %s \n", source, destination)
	err := Copy(source,destination, BUFFERSIZE)
	if err != nil{
		fmt.Printf("File copying failed : %q\n", err)
	}
}

func Copy(src, dst string, BUFFER_SIZE int64) error{
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcFileStat.Mode().IsRegular(){
		return fmt.Errorf("%s is not regular file\n", src)
	}

	source, err := os.Open(src)
	if err != nil{
		return err
	}
	defer source.Close()

	// dealing with the destination file
	_, err = os.Stat(dst)
	if err == nil{
		return fmt.Errorf("%s destination file already exits\n", dst)
	}
	destination, err := os.Create(dst)
	if err != nil{
		panic(err)
	}

	buf := make([]byte, BUFFER_SIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF{
			return err
		}
		if n == 0{
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
