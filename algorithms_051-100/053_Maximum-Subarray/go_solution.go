package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		nums []int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
		{nums: []int{1}},
		{nums: []int{5, 4, -1, 7, 8}},
		{nums: []int{1, 2, 3, 4}},
	}
	for _, c := range cases {
		out := maxSubArray(c.nums)
		fmt.Printf("Input: nums = %v\nOutput: %v\n\n", c.nums, out)
	}
}

// leetcode start

func maxSubArray(nums []int) int {
	theLength := len(nums)
	if theLength == 0 {
		return 0
	}
	dp := make([]int, theLength)
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < theLength; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
