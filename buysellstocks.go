package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		buy := prices[i]
		for j := i + 1; j < len(prices); j++ {
			sell := prices[j]
			res = max(res, sell-buy)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{10, 1, 5, 6, 7, 1}
	result := maxProfit(prices)
	fmt.Println(result) 
}
