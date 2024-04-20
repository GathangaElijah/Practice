package main

import (
	"fmt"

)

// This program takes an integer n and finds the sum of the digits. 
// If the sum of the digits is not a single digit, it recursively executes
// until the sum of the digits is a single digit
func DigitalRoot(n int) int {
	var eachDigit []int 
	var sum int
	for n > 0 {
		digits := n % 10
		eachDigit = append(eachDigit, digits)
		n = n/10
	}
	
	for _, digit := range eachDigit {
		sum += digit
	}
	if sum >= 10 {
		sum = DigitalRoot(sum)
	}

	return sum
}

func main() {
fmt.Println(DigitalRoot(3371672))
}