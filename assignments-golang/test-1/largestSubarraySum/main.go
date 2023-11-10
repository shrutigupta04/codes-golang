package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Find the subarray with largest sum")
	var n int
	fmt.Println("enter the size of array")
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scanln(&arr[i])
	}
	ans := maxSubArraySum(arr, n)
	fmt.Println(ans)

}

func maxSubArraySum(arr []int, N int) int {
	summ := 0
	maxii := math.MinInt64
	for i := 0; i < N; i++ {
		summ = summ + arr[i]
		maxii = max(summ, maxii)

		if summ < 0 {
			summ = 0
		}
	}
	return maxii
}
