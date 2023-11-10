package main

import "fmt"

func main() {
	fmt.Println("find the missing number in the array")
	fmt.Println("enter size of array")
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scanln(&arr[i])
	}
	ans := missingNumber(arr, n)
	fmt.Println(ans)
}
func missingNumber(arr []int, n int) int {
	sum := 0
	sum1 := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	for i := 0; i < n-1; i++ {
		sum1 = sum1 + arr[i]
	}
	return sum - sum1
}
