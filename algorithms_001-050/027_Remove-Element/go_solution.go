package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		nums []int
		val  int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
		{nums: []int{}, val: 3},
		{nums: []int{3}, val: 3},
		{nums: []int{3, 3}, val: 3},
	}

	for _, c := range cases {
		before := toJsonStr(c.nums)
		out := removeElement(c.nums, c.val)
		fmt.Printf("Input: nums = %s, val = %v \nOutput: %v, %s \n\n", before, c.val, out, toJsonStr(c.nums[:out]))
	}

}

func toJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// leetcode start

func removeElement(nums []int, val int) int {
	i, r := 0, len(nums)
	for i < r {
		if nums[i] == val {
			nums[i] = nums[r-1]
			r--
		} else {
			i++
		}
	}
	return r
}
