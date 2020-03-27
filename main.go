package main

import (
	"fmt"
)

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

func main() {
	//Test cases
	nums := []int{2, 7, 11, 15}
	fmt.Println("[0,1] = ", twoSum(nums, 9))

	nums2 := []int{8, 14, 1, 3, 2}
	fmt.Println("[1,3] = ", twoSum(nums2, 17))

	nums3 := []int{13, 7, 9, 11, 4}
	fmt.Println("[4,3] = ", twoSum(nums3, 15))

	nums4 := []int{3, 3}
	fmt.Println("[0,1] = ", twoSum(nums4, 6))
}

func twoSum(nums []int, target int) []int {
	indices := make([]int, 2)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				indices = []int{i, j}
			}
		}
	}
	return indices
}
