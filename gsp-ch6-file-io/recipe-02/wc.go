package main

import (
	"regexp"
	"flag"
	"os"
	"fmt"
	"bufio"
	"io"
)

func main(){
	wc := flag.Bool("w", false, "get word count")
	lc := flag.Bool("l", false, "get line count")
	cc := flag.Bool("c", false, "character count")
	flag.Parse()

	files := flag.Args()
	for _, file := range files {
		wCount := 0
		lCount := 0
		cCount := 0
		_, err := os.Stat(file)
		if err != nil {
			fmt.Printf("File could not be found: %s", file)
		}

		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error reading file: %s", file)
		}

		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading file: %s\n", err)
			}

			lCount++
			w, c := countWord(line)
			wCount += w
			cCount += c
		}
		switch {
			case *wc && *lc && *cc:
				fmt.Printf("%d %d %d  %s\n", cCount, wCount, lCount, file)
			case !*cc && *wc && *lc:
				fmt.Printf("%d %d %s\n", wCount, lCount, file)
			case *cc && !*wc && *lc:
				fmt.Printf("%d %d %s\n", cCount, lCount, file)
			case *cc && *wc && !*lc:
				fmt.Printf("%d %d %s\n", cCount, wCount, file)
			case *cc && !*wc && !*lc:
				fmt.Printf("%d %s\n", cCount, file)
			case !*cc && *wc && !*lc:
				fmt.Printf("%d %s\n", wCount , file)
			case !*cc && !*wc && *lc:
				fmt.Printf("%d %s\n", lCount , file)
			default:
				fmt.Printf("%d %d %d  %s\n", cCount, wCount, lCount, file)
		}
	}
}

func countWord(line string) (int, int){
	r := regexp.MustCompile("[^\\s]+")
	wCount := 0
	cCount := len(line)

	for range r.FindAllString(line,-1){
		wCount++
	}
	return wCount, cCount
}

