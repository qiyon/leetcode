package main

import "fmt"

func main() {
	cases := []struct {
		nums []int
		k    int
	}{
		{nums: []int{23, 2, 4, 6, 7}, k: 6},
		{nums: []int{23, 2, 6, 4, 7}, k: 6},
		{nums: []int{23, 2, 6, 4, 7}, k: 13},
		{nums: []int{5, 0, 0, 0}, k: 3},
	}

	for _, c := range cases {
		out := checkSubarraySum(c.nums, c.k)
		fmt.Printf("Input: nums = %v, k = %d\nOutput: %v\n\n", c.nums, c.k, out)
	}
}

// leetcode start

func checkSubarraySum(nums []int, k int) bool {
	iMap := make(map[int]int)
	iMap[0] = -1

	sum := 0
	for i, v := range nums {
		sum += v
		sum = sum % k
		if preIndex, ok := iMap[sum]; ok {
			if i-preIndex > 1 {
				return true
			}
		} else {
			iMap[sum] = i
		}
	}

	return false
}
