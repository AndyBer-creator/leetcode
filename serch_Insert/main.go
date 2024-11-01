// дан упорядоченный массив интов и значение,
// индекс которого нужно вернуть, если он есть в массиве.
// Если его нет в массиве вернуть индекс, на котором бы оно находилось
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, d := range nums {
		if d == target {
			return i
		} else if d > target {
			return i
		}
	}
	return len(nums)
}
func main() {
	nums := []int{1, 2, 4, 6, 7, 9}
	target := 8
	fmt.Println(searchInsert(nums, target))
}
