// URLify: Write a method to replace all spaces in a string with '%20'.
// You may assume that the string has sufficient space at the end to hold the additional characters,
// and that you are given the "true" length of the string.
// Input: 	"Mr John Smith    ", 13
// Output:	"Mr%20John%20Smith"
package main

import (
	"fmt"
	"unicode"
)

func main() {
	inputs := []string{
		"Mr John Smith     ",
	}

	for _, i := range inputs {
		fmt.Println(URLify(i, len(i)))
	}
}

func URLify(s string, i int) string {
	sByte := make([]rune, 0, len(s))

	bIndex := 0
	sIndex := 0
	for bIndex < len(s)-1 {
		c := rune(s[sIndex])

		if unicode.IsSpace(c) {
			sByte = append(sByte, []rune{'%', '2', '0'}...)
			bIndex += 3
		} else {
			sByte = append(sByte, c)
			bIndex++
		}
		sIndex++
	}

	return string(sByte)
}
