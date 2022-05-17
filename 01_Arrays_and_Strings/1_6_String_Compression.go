// Problem: Compress string using counts of repeated chars
// Example:
// aabcccccaaa -> a2b1c5a3
// abc -> abc
// abcc -> abc2 (len =4)  -> abcc
// Hints:
// #92: Compress string -> compare length
// #110: Not repeatedly concatenating strings together because not efficient

package main

import (
	"fmt"
	"strconv"
)

func compressString2(s string) string {
	c_s := ""
	count := 0
	char := s[0]
	for index := range s {
		if s[index] == char {
			count++
		} else {
			c_s += string(char) + strconv.Itoa(count)
			count = 1
			char = s[index]
		}
	}

	c_s += string(char) + strconv.Itoa(count)
	fmt.Printf("pre-compressed string:%s\n", c_s)
	if len(c_s) >= len(s) {
		return s
	}
	return c_s
}

func assignRepreatedChar(arr []byte, c_index int, char byte, count int) int {
	arr[c_index] = char
	c_index++
	count_str := strconv.Itoa(count)
	for index := range count_str {
		arr[c_index] = count_str[index]
		c_index++
	}
	return c_index
}

func compressString(s string) string {
	c_s := make([]byte, len(s)*2)
	c_index := 0
	count := 0
	char := s[0]
	for s_index := 0; s_index < len(s); s_index++ {
		if s[s_index] == char {
			count++
		} else {
			c_index = assignRepreatedChar(c_s, c_index, char, count)
			count = 1
			char = s[s_index]
		}
	}

	c_index = assignRepreatedChar(c_s, c_index, char, count)
	c_s = c_s[:c_index]
	fmt.Printf("pre-compressed string:%s\n", string(c_s))
	if len(c_s) >= len(s) {
		return s
	}
	return string(c_s)
}

func mainCompressString() {
	var s string
	s = "aabcccccaaa"
	fmt.Printf("Original string:%s\n", s)
	fmt.Printf("Compressed string:%s\n", compressString(s))
	s = "abcc"
	fmt.Printf("Original string:%s\n", s)
	fmt.Printf("Compressed string:%s\n", compressString(s))

	s = "aabbc"
	fmt.Printf("Original string:%s\n", s)
	fmt.Printf("Compressed string:%s\n", compressString(s))

	s = "aaaabbbcdddd"
	fmt.Printf("Original string:%s\n", s)
	fmt.Printf("Compressed string:%s\n", compressString(s))

}
