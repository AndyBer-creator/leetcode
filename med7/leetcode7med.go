// Given a signed 32-bit integer x, return x with its digits reversed.
//
//	If reversing x causes the value to go outside
//	the signed 32-bit integer range [-231, 231 - 1], then return 0.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))

}
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	res *= sign
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}
