package main

import (
	"fmt"
	"strings"
	"strconv"
	"unicode"
)

func main(){
	text := "Hi! Go is awesome"
	text = Indent(text,6)
	fmt.Println(text)

	text = UnIndent(text,3)
	fmt.Println(text)

	text = UnIndent(text,10)
	fmt.Println(text)

	text = IndentByRune(text, 10, '.')
	fmt.Println(text)
}

func IndentByRune(text string, indent int, r rune) string{
	return strings.Repeat(string(r),indent) + text
}

func Indent(text string, indent int) string{
	padding := indent + len(text)
	return fmt.Sprintf("% "+strconv.Itoa(padding)+"s", text)
}

func UnIndent(text string, indent int) string{
	count := 0
	for _, val := range text{
		if unicode.IsSpace(val){
			count++
		}
		if count == indent || !unicode.IsSpace(val){
			break
		}
	}
	return text[count:]
}
