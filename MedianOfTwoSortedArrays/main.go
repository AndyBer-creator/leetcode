package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	jointSlise := append(nums1, nums2...)
	fmt.Println(jointSlise)
	//посчитал балбес среднеарифметическое
	// sum := 0
	// for i := 0; i < len(jointSlise); i++ {
	// 	sum = sum + jointSlise[i]
	// 	fmt.Println(sum)
	// }
	// sumFloat := sum / 2
	// return float64(sumFloat)
	if len(jointSlise)%2 == 0 {
		return float64((jointSlise[len(jointSlise)/2] + jointSlise[len(jointSlise)/2-1]) / 2)
	}
	return float64((jointSlise[len(jointSlise)/2] + jointSlise[len(jointSlise)/2+1]) / 2)
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{2, 3}, []int{4, 5}))
}
