package main

import "fmt"

func maxArea(heights []int) int {
	res := 0
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			area := min(heights[i], heights[j]) * (j - i)
			if area > res {
				res = area
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 7, 2, 5, 4, 7, 3, 6}
	result := maxArea(height)
	fmt.Println(result)
}
