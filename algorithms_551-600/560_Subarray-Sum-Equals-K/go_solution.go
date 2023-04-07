package main

import "fmt"

func main() {
	cases := []struct {
		nums []int
		k    int
	}{
		{nums: []int{1, 1, 1}, k: 2},
		{nums: []int{1, 2, 3}, k: 3},
		{nums: []int{3, -2, 1, -1, 1, 1}, k: 0},
		{nums: []int{1, -1, 0}, k: 0},
	}

	for _, c := range cases {
		out := subarraySum(c.nums, c.k)
		fmt.Printf("Input: nums = %v, k = %d\nOutput: %d\n\n", c.nums, c.k, out)
	}
}

// leetcode start

func subarraySum(nums []int, k int) int {
	iMap := make(map[int]int)
	iMap[0] = 1

	out := 0
	sum := 0
	for _, v := range nums {
		sum += v
		if cnt, ok := iMap[sum-k]; ok {
			out += cnt
		}

		if _, ok := iMap[sum]; ok {
			iMap[sum] += 1
		} else {
			iMap[sum] = 1
		}
	}

	return out
}
