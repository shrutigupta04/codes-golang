package main

import "fmt"

func main() {
	fmt.Println("find if the strings are matching or not")

	var n int
	var m int
	fmt.Println("enter the size of string list")
	fmt.Scanln(&n)
	fmt.Println("enter the size of queries")
	fmt.Scanln(&m)
	stringList := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&stringList[i])
	}
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scanln(&queries[i])
	}
	fmt.Println(matchingStrings(stringList, queries))
}
func matchingStrings(stringList []string, queries []string) []int {
	c := 0
	var count []int
	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(stringList); j++ {
			if queries[i] == stringList[j] {
				c = c + 1
			}
		}
		count = append(count, c)
		c = 0
	}
	return count
}
