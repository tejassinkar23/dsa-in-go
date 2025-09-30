//go:build ignore

package main

import "fmt"

// func hasDuplicate(nums []int) bool {
// 	for i := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func hasDuplicate(nums []int) bool {
	seen := make(map[int]int)
	for i, num := range nums {
		if _, found := seen[num]; found {
			return true
		}
		seen[num] = i
	}
	return false
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 2, 3, 4}
	result1 := hasDuplicate(nums1)
	result2 := hasDuplicate(nums2)
	fmt.Println(result1)
	fmt.Println(result2)
}
