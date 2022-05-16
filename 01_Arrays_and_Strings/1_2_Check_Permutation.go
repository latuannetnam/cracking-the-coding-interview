// Problem: Give 2 strings, write a method to decide if one is a permutation of other
// Hints:
// #1 Permutation: ABC -> ACB, BAC, BCA, CAB, CBA
// #84 O(N Log N) or O(N)
// #122 Hashtable
// #131 Permutation = same size, different sort order
package main

import (
	"fmt"
	"sort"
	"strings"
)

func isStringPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	ss1 := strings.Split(s1, "")
	ss2 := strings.Split(s2, "")
	sort.Strings(ss1)
	sort.Strings(ss2)
	for index := range ss1 {
		if ss1[index] != ss2[index] {
			return false
		}
	}

	return true
}

func mainCheckPermutation() {
	s1 := "ABCA"
	s2 := "BCAA"
	fmt.Printf("s1: %s s2:%s\n", s1, s2)
	fmt.Printf("isPermutation: %t\n", isStringPermutation(s1, s2))
}
