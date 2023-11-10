package main

import (
	"fmt"
)

func main() {
	fmt.Println("check if the brackets are balanced or not")
	var s string
	fmt.Println("enter the string")
	fmt.Scanln(&s)
	fmt.Println(isBalanced(s))
}
func isBalanced(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}
	c := 0
	d := 0
	var stack []rune
	for i := 0; i < len(s); i++ {
		ch := rune(s[i])
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
			c = c + 1
		} else if len(stack) != 0 {
			d = d + 1
			ch2 := rune(stack[len(stack)-1])
			if ch == ')' && ch2 == '(' {
				stack = stack[0 : len(stack)-1]
			} else if ch == '}' && ch2 == '{' {
				stack = stack[0 : len(stack)-1]
			} else if ch == ']' && ch2 == '[' {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 && c != 0 && d != 0 && c+d == len(s) {
		return true
	} else {
		return false
	}

}
