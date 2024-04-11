// 1.2 Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.
package main

import "fmt"

func main() {
	inputs := map[[2]string]bool{
		{"elefante", "efelante"}: true,
		{"me", "met"}:            false,
		{"em", "emm"}:            false,
		{"emm", "em"}:            false,
		{"egg", "geg"}:           true,
	}

	for data, output := range inputs {
		if output != CheckPermutation(data[0], data[1]) {
			fmt.Printf("Wrong output, expecting %t, for strings %s %s\n", output, data[0], data[1])
			return
		}
	}

	fmt.Println("All cases were successful")
}

func CheckPermutation(s1, s2 string) bool {
	mapS1 := make(map[rune]int)

	for _, cS1 := range s1 {
		mapS1[cS1]++
	}

	for _, cS2 := range s2 {
		if mapS1[cS2] <= 0 {
			return false
		}

		mapS1[cS2]--
		if mapS1[cS2] <= 0 {
			delete(mapS1, cS2)
		}
	}

	return len(mapS1) == 0
}
