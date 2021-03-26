package main

import (
	"fmt"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] = i
	}
	for i := range nums {
		impl := target - nums[i]
		if index, ok := m[impl]; ok && index != i {
			return []int{i, index}
		}
	}
	return nil
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Print(result)
}
