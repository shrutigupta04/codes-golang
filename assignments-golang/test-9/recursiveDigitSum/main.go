package main

import (
	"fmt"
)

func main() {
	fmt.Println("find the recursive digit sum of a number")
	var n int
	fmt.Println("enter number")
	fmt.Scanln(&n)
	var k int
	fmt.Println("enter recursive number")
	fmt.Scanln(&k)
	ans := superDigit(n, k)
	fmt.Println(ans)
}
func summ(s int, sum2 int) int {
	for s != 0 {
		rem := s % 10
		sum2 = sum2 + rem
		s = int(s / 10)
	}
	if sum2/10 == 0 {
		return sum2
	} else {
		s = sum2
		sum2 = 0
	}
	return summ(s, sum2)
}
func superDigit(n int, k int) int {
	s := 0
	num := int(n)
	for num != 0 {
		s = s + (num % 10)
		num = int(num / 10)
	}
	sum2 := 0
	s = s * k
	sum1 := summ(s, sum2)
	return sum1
}
