// Problem: medthod isSubstring return one word is a substring of another word. Use isSubstring to check if s2 if a rotation of s1
// Example: waterbottle -> erbottewat
// Hint:
// #34: if string is a rotation of another, then it's a rotation at particular point: waterbottle -> [wat]erbottle -> errbottle[wat]
// #88:s1 = xy => s2 = yx => yx + yx = y[xy]x
// #104: try s2 + s2 => s1: terbottlewa + terbottlewa -> terbottle[waterbottle]wa
package main

import (
	"fmt"
	"strings"
)

func IsRotation(s1, s2 string) bool {
	if len(s2) != len(s1) {
		return false
	}
	s3 := s2 + s2
	return strings.Contains(s3, s1)

}

func mainStringRotation() {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"waterbottle", "terbottlewa", true},
		{"hellomynameis", "nameishellomy", true},
		{"waterbottle", "water", false},
		{"waterbottle", "elttobretaw", false},
		{" ", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		actual := IsRotation(c.input1, c.input2)
		fmt.Printf("Inputs %s, %s isRotation:%t\n", c.input1, c.input2, actual)
		if actual != c.expected {
			fmt.Printf("Inputs %s, %s. Expected: %t, actual: %t\n",
				c.input1, c.input2, c.expected, actual)
		} else {
			fmt.Println("OK!")
		}
	}
}
