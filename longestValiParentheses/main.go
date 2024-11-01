package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	stack := []int{-1} // use -1 to handle edge case of valid substring at start
	maxLength := 0

	for i, char := range s {
		if char == '(' {
			stack = append(stack, i) // push the index of '(' onto the stack
		} else {
			stack = stack[:len(stack)-1] // pop the last index

			if len(stack) == 0 {
				stack = append(stack, i) // found an unmatched ')' so reset the base index
			} else {
				// calculate the length of the current valid parentheses substring
				validLength := i - stack[len(stack)-1]
				if validLength > maxLength {
					maxLength = validLength
				}
			}
		}
	}

	return maxLength
}

func main() {
	// Test cases
	fmt.Println(longestValidParentheses("(()"))      // Output: 2
	fmt.Println(longestValidParentheses(")()())"))   // Output: 4
	fmt.Println(longestValidParentheses(""))         // Output: 0
	fmt.Println(longestValidParentheses("()(()"))    // Output: 2
	fmt.Println(longestValidParentheses("()()"))     // Output: 4
	fmt.Println(longestValidParentheses("((())"))    // Output: 4
	fmt.Println(longestValidParentheses("()((()))")) // Output: 8
}
