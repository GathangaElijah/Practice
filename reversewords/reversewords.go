package main

import (
	"fmt"
)

// This function takes an input of type string and reverses it.
// The method showed here
func ReverseWords(str string) string {
	var newstring string
	isSpace := false
	for _, char := range str {
		if char == ' ' && isSpace {
			isSpace = true
			newstring = ""
		} else if !isSpace {
			newstring = string(char) + newstring
		}
	}
	isSpace = false

	return string(newstring)
}

func main() {
	str := "The quick"
	fmt.Println(ReverseWords(str))
}
