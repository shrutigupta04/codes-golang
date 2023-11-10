package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("find the power sum of a number")
	var x float64
	var n float64
	fmt.Println("enter total")
	fmt.Scanln(&x)
	fmt.Println("enter power")
	fmt.Scanln(&n)
	ans := powerSum(x, n, 1)
	fmt.Println(ans)
}
func powerSum(total float64, power float64, num float64) int {
	val := total - math.Pow(num, power)
	if val == 0 {
		return 1
	}
	if val < 0 {
		return 0
	}
	return powerSum(val, power, num+1) + powerSum(total, power, num+1)

}
