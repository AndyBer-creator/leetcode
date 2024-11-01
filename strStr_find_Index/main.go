// Given two strings needle and haystack, return the index of the first
//occurrence of needle in haystack, or -1 if needle is not part of haystack.
// Example 1:
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.
// Example 2:
// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.

package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0 
	}

	index := strings.Index(haystack, needle)
	return index
}

func main() {
	// Test cases
	fmt.Println(strStr("sadbutsad", "sad"))  // Output: 0
	fmt.Println(strStr("leetcode", "leeto")) // Output: -1
}
