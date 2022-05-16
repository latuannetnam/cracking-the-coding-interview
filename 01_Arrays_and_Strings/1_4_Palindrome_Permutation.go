// Problem: given a string, check if string is a permutation of palindrome
// Example:
//  Tact Coa -> "taco cat", "tatco cta"
package main

import "fmt"

func isStringPermutationPalindrome(s string) bool {
	hash := map[string]int{}
	// Build map to count ocurrents of char in string
	odd_count := 0
	for index := range s {
		char := string(s[index])
		_, ok := hash[char]
		if ok {
			hash[char] = hash[char] + 1
		} else {
			hash[char] = 1
		}

		if hash[char]%2 == 1 {
			odd_count++
		} else {
			odd_count--
		}
	}

	// Check if string is able to be palindrome

	// for index := range hash {
	// 	if hash[index]%2 == 1 {
	// 		odd_count++
	// 	}
	// }

	return odd_count <= 1
}

func mainCheckStringPermutationPalindrome() {
	s := "tactcoa" //palindrome: tacocat

	fmt.Printf("string:%s isPermutation:%t\n", s, isStringPermutationPalindrome(s))
}
