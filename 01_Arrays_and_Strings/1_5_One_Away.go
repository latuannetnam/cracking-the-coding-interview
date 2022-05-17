// Problem: Check string A is a edit of string B
// Example:
// pale -> ple (remove a)
// pales -> pale (remove s)
// pale -> palek (insert k)
// pale -> bale (replace p -> b)
// pale -> bake (replace 2 chars -> false)
// pale -> pak (replace 2 chars -> false)

// Hint:
// #23: check each condition separately
// #97: relationship between insert char vs remove char
// #130: do 3 checks in one pass

package main

import (
	"fmt"
)

func insertChar(s string, index int, char string) string {
	return s[:index] + char + s[index:]
}

func removeChar(s string, index int) string {
	if index < len(s)-1 {
		return s[:index] + s[index+1:]
	} else {
		return s[:index]
	}

}

func isEditString2(s1 string, s2 string) bool {
	s1_len := len(s1)
	s2_len := len(s2)

	// s1 len and s2 len must not diffent by more than 1 char
	if (s1_len-s2_len)*(s1_len-s2_len) > 1 {
		return false
	} else if s1_len == s2_len {
		count := 0
		for index := 0; index < s1_len; index++ {
			if s1[index] != s2[index] {
				count++
			}
			if count > 1 {
				return false
			}
		}
		if count == 0 {
			return false
		}

	} else {
		// Find 1st different char index
		index := 0
		for index = 0; index < s2_len && index < s1_len; index++ {
			if s1[index] != s2[index] {
				break
			}
		}

		// Insert different char to make 2 strings  equal
		if s1_len > s2_len {
			s2 = insertChar(s2, index, string(s1[index]))
		} else {
			s1 = insertChar(s1, index, string(s2[index]))
		}

		fmt.Printf("s1:%s s2 new:%s\n", s1, s2)
		return s1 == s2
	}

	return true
}

func isEditString(s1 string, s2 string) bool {
	index := 0
	// Assume s1_len >= s2_len
	for index = 0; index < len(s2) && index < len(s1); index++ {
		if s1[index] != s2[index] {
			break
		}
	}

	// Insert different char to make 2 strings  equal
	if len(s1) > len(s2) {
		s2 = insertChar(s2, index, string(s1[index]))
	} else if len(s1) < len(s2) {
		s1 = insertChar(s1, index, string(s2[index]))
	} else {
		// remove different char index to make 2 strings  equal
		s1 = removeChar(s1, index)
		s2 = removeChar(s2, index)
	}

	fmt.Printf("s1:%s s2 new:%s\n", s1, s2)
	return s1 == s2
}

// Implement solution in book
func isOneAway(s1 string, s2 string) bool {
	// s1 len and s2 len must not diffent by more than 1 char
	s1_len := len(s1)
	s2_len := len(s2)
	if (s1_len-s2_len)*(s1_len-s2_len) > 1 {
		return false
	}
	// Make sure len(s1) < len(s2)
	if len(s1) > len(s2) {
		// Swap s1 <=> s2
		s1, s2 = s2, s1
	}
	index1 := 0
	index2 := 0
	isDifferent := false
	for index1 < len(s1) && index2 < len(s2) {
		if s1[index1] != s2[index2] {
			// Allow only 1 different
			if isDifferent {
				return false
			}
			isDifferent = true
			if s1_len == s2_len {
				// move short index
				index1++
			}
		} else {
			// move short index
			index1++
		}
		// move longer index
		index2++
	}

	return true

}

func mainCheckEditString() {
	s1 := "pales"
	s2 := "pale"
	fmt.Printf("s1: %s s2:%s\n", s1, s2)
	fmt.Printf("isEdit:%t isOneAway:%t\n", isEditString(s1, s2), isOneAway(s1, s2))

	s1 = "pale"
	s2 = "pake"
	fmt.Printf("s1: %s s2:%s\n", s1, s2)
	fmt.Printf("isEdit:%t isOneAway:%t\n", isEditString(s1, s2), isOneAway(s1, s2))

	s1 = "pale"
	s2 = "pales"
	fmt.Printf("s1: %s s2:%s\n", s1, s2)
	fmt.Printf("isEdit:%t isOneAway:%t\n", isEditString(s1, s2), isOneAway(s1, s2))

	s1 = "pale"
	s2 = "bake"
	fmt.Printf("s1: %s s2:%s\n", s1, s2)
	fmt.Printf("isEdit:%t isOneAway:%t\n", isEditString(s1, s2), isOneAway(s1, s2))

	s1 = "paless"
	s2 = "pale"
	fmt.Printf("s1: %s s2:%s\n", s1, s2)
	fmt.Printf("isEdit:%t isOneAway:%t\n", isEditString(s1, s2), isOneAway(s1, s2))

}
