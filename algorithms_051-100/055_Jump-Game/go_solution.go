package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		nums []int
	}{
		{nums: []int{2, 3, 1, 1, 4}},
		{nums: []int{3, 2, 1, 0, 4}},
	}
	for _, c := range cases {
		fmt.Printf("Input: nums= %v\nOutput:%v\n\n", c.nums, canJump(c.nums))
	}
}

// leetcode start

func canJump(nums []int) bool {
	theLength := len(nums)
	if theLength <= 1 {
		return true
	}
	last := theLength - 1
	for i := theLength - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last <= 0
}
