package main

import (
	"fmt"
	"strings"
)

// This function takes an input of type string and reverses it.
// The method showed here
func ReverseWords(str string) string {
	var newstring string
	//var revstring []string
	substrings := strings.Fields(str)
	for _, word := range substrings {
		for i := len(word) - 1; i >= 0; i-- {
			newstring = newstring + string(word[i]) 
		}
	}
	//newstring = strings.Join(revstring, " ")
	return newstring
}

func main() {
	str := "The quick"
	fmt.Println(ReverseWords(str))
}
