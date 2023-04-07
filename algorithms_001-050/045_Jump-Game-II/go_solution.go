package main

import (
	"fmt"
)

func main() {
	cases := [][]int{
		{10, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 1},
		{0},
		{3, 2, 3, 4, 2, 5, 1},
	}
	for _, c := range cases {
		fmt.Printf("Input: nums = %v\nOutput: %d\n\n", c, jump(c))
	}
}

// leetcode start

func jump(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return 0
	}

	minOf := make([]int, numsLen)
	minOf[0] = 0
	hasRecordsIndex := 0
	for i := 0; i < numsLen; i++ {
		if i > hasRecordsIndex {
			continue
		}

		iJump := nums[i]
		minStep2I := minOf[i]
		for k := i + 1; k <= (i+iJump) && k < numsLen; k++ {
			if k > hasRecordsIndex {
				minOf[k] = minStep2I + 1
				hasRecordsIndex++
			} else {
				minOf[k] = min(minOf[k], minStep2I+1)
			}
		}
	}
	return minOf[numsLen-1]
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
