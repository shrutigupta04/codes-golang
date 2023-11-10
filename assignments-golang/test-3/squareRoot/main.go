package main

import "fmt"

func main() {
	fmt.Println("print the absolute square root of a number")
	var x int
	fmt.Println("enter the number")
	fmt.Scanln(&x)
	for i := 1; i <= x; i++ {
		if i*i > x {
			fmt.Println(i - 1)
			break
		}
	}
}
