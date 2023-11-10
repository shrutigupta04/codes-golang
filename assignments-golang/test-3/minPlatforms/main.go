package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("find the min number of platforms required in a railway station")
	var n int
	fmt.Println("enter size")
	fmt.Scanln(&n)
	arrival := make([]int, n)
	departute := make([]int, n)
	for i := range arrival {
		fmt.Scanln(&arrival[i])
	}
	for i := range departute {
		fmt.Scanln(&departute[i])
	}
	ans := minPlatform(n, arrival, departute)
	fmt.Println(ans)
}

func minPlatform(n int, arr []int, dep []int) int {
	sort.Ints(arr)
	sort.Ints(dep)
	i := 0
	j := 0
	count := 0
	var l []int
	for i < n && j < n {
		if arr[i] <= dep[j] {
			count = count + 1
			l = append(l, count)
			i++
		} else {
			count = count - 1
			l = append(l, count)
			j = j + 1
		}
	}
	sort.Ints(l)
	return l[len(l)-1]

}
