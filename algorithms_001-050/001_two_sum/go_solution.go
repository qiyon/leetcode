package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		nums   []int
		target int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9},
		{nums: []int{3, 2, 4}, target: 6},
		{nums: []int{3, 3}, target: 6},
	}

	for _, c := range cases {
		out := twoSum(c.nums, c.target)
		fmt.Printf("nums:%s target:%d \noutput:%s \n", toJsonStr(c.nums), c.target, toJsonStr(out))
	}
}

func toJsonStr(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

// leetcode start

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}

	for i, v := range nums {
		if preIndex, ok := hashMap[v]; ok {
			return []int{preIndex, i}
		}
		hashMap[target-v] = i
	}

	return []int{}
}
