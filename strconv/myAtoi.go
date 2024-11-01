package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
	fmt.Println(myAtoi("1234rfs"))
}
func myAtoi(s string) int {
	result := 0
	sign := 1
	index := 0
	n := len(s)

	// Skip leading whitespace
	for index < n && unicode.IsSpace(rune(s[index])) {
		index++
	}

	// Check for sign
	if index < n && (s[index] == '+' || s[index] == '-') {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}

	// Read digits
	for index < n && s[index] >= '0' && s[index] <= '9' {
		digit := int(s[index] - '0')

		// Check for overflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		index++
	}

	return result * sign
}
// func myAtoi(s string) int {
//     var c string
//     var a, d, cnt int = 0, 1, 0

//     for _,i:=range s {
//         if i==' ' {
//             if cnt>0 {
//                 break
//             }
//             continue
//         }
//         cnt++
//         if i=='-' || i=='+' {
//             if len(c)>0 {
//                 break
//             }
//             a++
//             if a>1 {
//                 break
//             }
//             if i=='-' {
//                 d=-1
//             }
//         } else if i>='0' && i<='9' {
//             c+=string(i)
//         } else {
//             break
//         }
//     }

//     n:=0
//     for i:=0; i<len(c); i++ {
//         n*=10
//         n+=int(c[i])-48
//         if n*d > ((1<<31)-1) {
//             return ((1<<31)-1)
//         }
//         if n*d < -(1<<31) {
//             return -(1<<31)
//         }
//     }
    
//     return n*d
// }