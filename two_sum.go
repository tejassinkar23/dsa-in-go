package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, found := seen[diff]; found {
			return []int{j, i}
		}
		seen[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{3, 4, 5, 6}
	target := 7
	result := twoSum(nums, target)
	fmt.Println(result)
}
