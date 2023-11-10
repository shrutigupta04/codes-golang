package main

import "fmt"

func main() {
	fmt.Println("find if a given string is palindrome")
	var s string
	fmt.Println("enter the string")
	fmt.Scanln(&s)
	ans := isPalindrome(s)
	fmt.Println(ans)
}

func isPalindrome(s string) int {
	for i := 0; i < int(len(s)/2); i++ {
		if s[i] != s[len(s)-i-1] {
			return 0
		}
	}
	return 1
}
