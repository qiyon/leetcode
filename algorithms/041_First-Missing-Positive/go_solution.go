package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		nums []int
	}{
		{nums: []int{3, 4, -1, 1}},
		{nums: []int{1, 2, 0}},
		{nums: []int{1, 2, 0, 2, 2, -2, 2, 2, 2, 4, 5, 6, 7, 8}},
		{nums: []int{1, 2, 3}},
	}

	for _, c := range cases {
		out := firstMissingPositive(c.nums)
		fmt.Printf("Nums= %v \nOutput= %v \n\n", c.nums, out)
	}
}

// leetcode start

func firstMissingPositive(nums []int) int {
	// make nums, nums[i] == i+1
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1 {
			toPosition := nums[i] - 1
			if nums[toPosition] == toPosition+1 {
				break
			}
			nums[i], nums[toPosition] = nums[toPosition], nums[i]
		}
	}
	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}
	return len(nums) + 1
}
