package main

import "fmt"

func main() {
	fmt.Println("merge two sorted arrays")
	var m int
	var n int
	fmt.Println("length of first array")
	fmt.Scanln(&m)
	fmt.Println("length of second array")
	fmt.Scanln(&n)
	ar1 := make([]int, m)
	ar2 := make([]int, n)
	fmt.Println("enter first array: ")
	for i := 0; i < m; i++ {
		fmt.Scanln(&ar1[i])
	}
	fmt.Println("enter second array: ")
	for i := 0; i < n; i++ {
		fmt.Scanln(&ar2[i])
	}
	merge(ar1, ar2, &m, &n)
	fmt.Println("the first array becomes: ")
	for i := range ar1 {
		fmt.Println(ar1[i])
	}
	fmt.Println("the second array becomes: ")
	for i := range ar2 {
		fmt.Println(ar2[i])
	}
}
func merge(ar1 []int, ar2 []int, m *int, n *int) {
	for i := *n - 1; i >= 0; i-- {
		last := ar1[*m-1]
		j := *m - 2
		for j >= 0 && ar1[j] > ar2[i] {
			ar1[j+1] = ar1[j]
			j = j - 1
		}
		if last > ar2[i] {
			ar1[j+1] = ar2[i]
			ar2[i] = last
		}
	}
}
