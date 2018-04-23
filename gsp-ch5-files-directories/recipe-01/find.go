package main

import (
	"os"
	"fmt"
	"flag"
	"path/filepath"
	"strings"
	"regexp"
)

func main(){
	socket := flag.Bool("s", false, "search file which are of socket type")
	pipes := flag.Bool("p", false, "search file which are of pipeline type")
	files := flag.Bool("f", false, "search simple file types")
	dir := flag.Bool("d", false, "search file which are of directory type")
	links := flag.Bool("sl", false, "search file synbolic link type")

	exFileName := flag.String("x", "", "Exclude file name")
	ext := flag.String("ext", "", "exlude all files with this extension")
	re := flag.String("re", "", "match the files with the regular expression")

	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0{
		fmt.Println("Not enough arguments !")
		os.Exit(1)
	}
	Path := flags[0]
	printAll := false
	if !(*socket || *pipes || *links || *dir || *files) {
		printAll = true
	}

	if *socket && *pipes && *links && *dir && *files {
		printAll = true
	}

	walk := func (path string, info os.FileInfo, err error) error{
		fileInfo, err := os.Stat(path)
		if err != nil{
			return err
		}

		if excludeNames(path, *exFileName){
			return nil
		}
		if excludeExtension(path, *ext){
			return nil
		}

		if regularExpressionMatch(path, *re) == false{
			return nil
		}

		if printAll{
			fmt.Println(path)
			return nil
		}
		mode := fileInfo.Mode()
		if mode.IsRegular() && *dir{
			fmt.Println(path)
			return nil
		}
		fileInfo, _ = os.Lstat(path)
		if fileInfo.Mode()&os.ModeSymlink != 0{
			if *links {
				fmt.Println(path)
				return nil
			}
		}
		if fileInfo.Mode()&os.ModeNamedPipe != 0{
			if *pipes {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeSocket != 0{
			if *socket {
				fmt.Println(path)
				return nil
			}
		}
		return nil
	}

	err := filepath.Walk(Path, walk)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func excludeNames(name string, excludeFile string) bool{
	if excludeFile == ""{
		return false
	}
	if filepath.Base(name) == excludeFile{
		return true
	}
	return false
}

func excludeExtension(name string, excludeExt string) bool{
	if excludeExt == ""{
		return false
	}

	basename := filepath.Base(name)
	s := strings.Split(basename, ".")
	baseExt := s[len(s)-1]
	if baseExt == excludeExt{
		return true
	}
	return false
}

func regularExpressionMatch(name string, expression string) bool{
	if expression == ""{
		return true
	}
	r, err := regexp.Compile(expression)
	if err != nil{
		//fmt.Println("Error regular express ", err)
		return true
	}
	matched := r.MatchString(name)
	return matched
}

