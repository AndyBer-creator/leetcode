package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	result := []string{}
	generate(&result, "", n, n)
	return result
}


func generate(result *[]string, current string, open int, close int) {
		if open == 0 && close == 0 {
		*result = append(*result, current)
		return
	}
	if open > 0 {
		generate(result, current+"(", open-1, close)
	}
	if close > open {
		generate(result, current+")", open, close-1)
	}
}

func main() {
	n1 := 3
	fmt.Println("Input:", n1)
	fmt.Println("Output:", generateParenthesis(n1))

	n2 := 1
	fmt.Println("\nInput:", n2)
	fmt.Println("Output:", generateParenthesis(n2))
}
