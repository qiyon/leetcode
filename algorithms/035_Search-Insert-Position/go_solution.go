package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6, 7, 8, 10}
	cases := []struct {
		nums   []int
		target int
	}{
		{nums: nums, target: 0},
		{nums: nums, target: 2},
		{nums: nums, target: 5},
		{nums: nums, target: 9},
		{nums: nums, target: 11},
	}

	for _, c := range cases {
		out := searchInsert(c.nums, c.target)
		fmt.Printf("nums: %v, target: %v \nout: %v \n\n", c.nums, c.target, out)
	}
}

//leetcode begin

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	if end < 0 {
		return 0
	}
	if nums[0] > target {
		return 0
	}
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	if nums[end] > target {
		return start + 1
	} else {
		return end + 1
	}
}
