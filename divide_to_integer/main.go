package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend/divisor == 2147483648 {
		return dividend/divisor - 1
	}
	return dividend / divisor
}
func main() {
	fmt.Println(divide(12456, 44))
}
