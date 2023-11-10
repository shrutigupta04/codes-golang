package main

import "fmt"

func main() {
	fmt.Println("find the longest increasing subsequence")
	var n int
	fmt.Println("enter size of array")
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&nums[i])
	}
	fmt.Println(longIncSubseq(nums))
}
func longIncSubseq(nums []int) int {
	var piles [][]int
	for i := range nums {
		left := 0
		right := len(piles) - 1
		for left <= right {
			mid := (left + right) / 2
			if piles[mid][-1] < i {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left == len(piles) {
			var digit []int
			digit = append(digit, i)
			piles = append(piles, digit)
		} else {
			piles[left] = append(piles[left], i)
		}
	}
	for i := 0; i < len(piles); i++ {
		fmt.Println(piles[i])
	}
	return len(piles)
}
