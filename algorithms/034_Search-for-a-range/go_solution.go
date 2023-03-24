package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 1, 2, 3, 3, 4, 5, 5, 9, 10, 10, 10, 12, 15, 15, 16}
	cases := []struct {
		nums   []int
		target int
		should []int
	}{
		{nums: nums, target: 1, should: []int{0, 1}},
		{nums: nums, target: 20, should: []int{-1, -1}},
		{nums: []int{2, 2}, target: 2, should: []int{0, 1}},
	}

	for _, c := range cases {
		out := searchRange(c.nums, c.target)
		fmt.Printf("nums: %v, target: %v \nshould: %v, out: %v \n\n", c.nums, c.target, c.should, out)
	}
}

//leetcode start

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := findL(nums, target)
	r := -1
	if l != -1 {
		r = findR(nums, target, l)
	}
	return []int{l, r}
}

func findL(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	start := 1
	end := len(nums) - 1
	if end <= 0 {
		return -1
	}
	for start < end-1 {
		//fmt.Printf("L %v %v \n", start, end)
		mid := (start + end) / 2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid
		}
	}
	if isTargetLeft(nums, target, start) {
		return start
	}
	if isTargetLeft(nums, target, end) {
		return end
	}
	return -1
}

func isTargetLeft(nums []int, target int, index int) bool {
	return nums[index] == target && nums[index-1] < target
}

func findR(nums []int, target int, start int) int {
	length := len(nums)
	if nums[length-1] == target {
		return length - 1
	}
	end := length - 2
	for start < end-1 {
		//fmt.Printf("R %v %v \n", start, end)
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid
		} else {
			end = mid
		}
	}
	if isTargetRight(nums, target, start) {
		return start
	}
	if isTargetRight(nums, target, end) {
		return end
	}
	return -1
}

func isTargetRight(nums []int, target int, index int) bool {
	return nums[index] == target && nums[index+1] > target
}
