// 1.1 Is Unique: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?.
package main

import "fmt"

func main() {
	inputs := map[string]bool{
		"input_ok":          true,
		"wrong_input_value": false,
	}

	for text, expectedOutput := range inputs {
		if expectedOutput != IsUnique(text) {
			fmt.Printf("The output is invalid for %s, expecting %t\n", text, expectedOutput)
		}
	}

	fmt.Println("All cases succeed")
}

func IsUnique(s string) bool {
	data := make(map[rune]bool)

	for _, c := range s {
		if data[c] {
			return false
		}

		data[c] = true
	}

	return true
}
