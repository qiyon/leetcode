package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		nums []int
	}{
		{nums: []int{1, 1, 2}},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
		{nums: []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3}},
	}

	for _, c := range cases {
		before := toJsonStr(c.nums)
		out := removeDuplicates(c.nums)
		fmt.Printf("Input: nums = %s \nOutput: %d, nums = %s\n\n", before, out, toJsonStr(c.nums))
	}
}

func toJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// leetcode start

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	l, i, j, r := 0, 1, len(nums)-2, len(nums)-1
	for i <= j {
		for i <= j && nums[i] == nums[l] {
			i++
		}
		if i <= j && nums[i] != nums[l] {
			nums[l+1], nums[i] = nums[i], nums[l+1]
			l++
			i++
		}
		for i <= j && nums[j] == nums[r] {
			j--
		}
		if i <= j && nums[j] != nums[r] {
			nums[r-1], nums[j] = nums[j], nums[r-1]
			r--
			j--
		}
	}

	if nums[l] == nums[r] {
		r++
	}
	for m := r; m < len(nums); m++ {
		nums[l+1+m-r] = nums[m]
	}
	return l + 1 + len(nums) - 1 - r + 1
}
