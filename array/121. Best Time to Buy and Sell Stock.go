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
	var maxProfit int
	minPrice := math.MaxInt32
	for _, v := range prices {
		curProfit := max(v-minPrice, 0)
		maxProfit = max(maxProfit, curProfit)
		minPrice = min(minPrice, v)
	}
	return maxProfit
}

// func main() {
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
// }
