//Create a function taking a positive integer between 1 and 3999
//(both included) as its parameter and returning a string
//containing the Roman Numeral representation of that integer.
//Modern Roman numerals are written by expressing each digit
//separately starting with the leftmost digit and
//skipping any digit with a value of zero.
//There cannot be more than 3 identical symbols in a row.
package main

import "fmt"

func Solution(number int) string {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	if number < 1 || number > 3999 {
		return "error: number must be between 1 and 3999\n"
	}
	var result = ""
	for num := number; num > 0; num /= 10 {
		lastNumber := num % 10
		if lastNumber != 0 {
			present, ok := romanMap[lastNumber]
			if ok {
				result = present + result
			}
		}
	}

	return result
}

func main() {
	want := 3991
	fmt.Printf("The Roman Value is %v\n", Solution(want))
}
