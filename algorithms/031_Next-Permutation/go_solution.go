package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		nums []int
	}{
		{nums: []int{1, 2, 3}},
		{nums: []int{3, 2, 1}},
		{nums: []int{1, 1, 5}},
		{nums: []int{1, 2, 3, 4}},
	}

	for _, c := range cases {
		fmt.Printf("Input: nums = %v\n", c.nums)
		nextPermutation(c.nums)
		fmt.Printf("Output: %v \n\n", c.nums)
	}
}

// leetcode start

func nextPermutation(nums []int) {
	cnt := len(nums)
	if cnt <= 1 {
		return
	}

	plode := 0
	hiil := 0
	for i := cnt - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			plode = i - 1
			hiil = nums[plode]
			break
		}
	}
	exchangeIdx := plode + 1
	for i := plode + 1; i < cnt; i++ {
		if nums[i] > hiil && nums[i] <= nums[exchangeIdx] {
			exchangeIdx = i
		}
	}

	//swap
	nums[plode], nums[exchangeIdx] = nums[exchangeIdx], nums[plode]
	intSort(nums, plode+1, cnt-1)
}

func intSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	mid, i := nums[left], left+1
	head, tail := left, right
	for i <= tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	nums[head] = mid
	intSort(nums, left, head-1)
	intSort(nums, head+1, right)
}
