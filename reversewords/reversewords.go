package main

import (
	"fmt"
	"strings"
)

// This function takes an input of type string and reverses it.
// The method showed here swaps the indexex of the substrings.
func ReverseWords(str string) string {
	var reversedStr []string
	var finalStr string
	subStr := strings.Fields(str)

	for _, str := range subStr {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		reversedStr = append(reversedStr, string(runes))
		finalStr = strings.Join(reversedStr, " ")
	}
	return finalStr
}

func main() {
	str := "The quick"
	fmt.Println(ReverseWords(str))
}
