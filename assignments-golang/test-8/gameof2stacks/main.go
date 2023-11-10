package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("play this game of 2 stacks")
	var lst []int
	var n int
	fmt.Println("enter no. of elements in list")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&lst[i])
	}
	var lst1 []int
	var m int
	fmt.Println("enter no. of elements in list2")
	fmt.Scanln(&m)
	for i := 0; i < m; i++ {
		fmt.Scanln(&lst1[i])
	}
	var maxs int
	fmt.Println("enter the maximum sum")
	fmt.Scanln(&maxs)
	ans := twoStacks(maxs, lst, lst1)
	fmt.Println(ans)
}
func twoStacks(maxSum int, a []int, b []int) int {
	sum := 0
	st1 := 0
	st2 := 0
	var arr []int
	for i := 0; i < len(a); i++ {
		if (sum + a[i]) > maxSum {
			break
		}
		sum = sum + a[i]
		st1 = st1 + 1
		arr = append(arr, st1)
	}
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
		st2++
		for sum > maxSum && st1 > 0 {
			sum = sum - a[st1-1]
			st1 = st1 - 1
		}
		if sum <= maxSum {
			arr = append(arr, st1+st2)
		}

	}
	sort.Ints(arr)
	return arr[len(arr)-1]
}
