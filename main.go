package main

import (
	"fmt"

)

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