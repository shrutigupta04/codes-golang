package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("maximize the profit in stocks")
	var n int
	fmt.Println("enter the size of prices to be bought of stocks")
	fmt.Scanln(&n)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&prices[i])
	}
	fmt.Println(stockMax(prices))
}
func stockMax(prices []int) int {
	maxprice := math.MinInt64
	profit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxprice {
			maxprice = prices[i]
		}
		if maxprice > prices[i] {
			profit = profit + maxprice - prices[i]
		}
	}
	return profit

}
