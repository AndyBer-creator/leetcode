package main

import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// The maximum possible length of the result
	result := make([]int, len(num1)+len(num2))

	// Reverse traverse both strings
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			// Multiply the digits
			mul := (int(num1[i] - '0')) * (int(num2[j] - '0'))
			// Position in the result array
			p1 := i + j
			p2 := i + j + 1

			// Add to the result position
			total := mul + result[p2]
			result[p2] = total % 10  // Current digit
			result[p1] += total / 10 // Carry
		}
	}

	// Convert result array to string, skipping leading zeros
	var sb strings.Builder
	for _, v := range result {
		if !(sb.Len() == 0 && v == 0) { // Skip leading zeros
			sb.WriteByte(byte(v + '0'))
		}
	}

	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}

func main() {
	// Example usage
	fmt.Println(multiply("123", "456")) // Output: "56088"
	fmt.Println(multiply("2", "3"))     // Output: "6"
}
