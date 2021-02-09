package main

import (
	"math"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	maxProfit, minPrice := 0, math.MaxInt32
	for _, v := range prices {
		maxProfit = max(maxProfit, v-minPrice)
		minPrice = min(minPrice, v)
	}
	return maxProfit
}

// func main() {
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
// }
