package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct{ nums []int }{
		{nums: []int{-1, 0, 1, 2, -1, -4}},
		{nums: []int{0, 1, 1}},
		{nums: []int{0, 0, 0}},
		{nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
		{nums: []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}},
		{nums: []int{-1, 0, 0, 0, 1}},
		{nums: []int{0, 0, 0, 1}},
		{nums: []int{-1, 0, 0, 0}},
		{nums: []int{0, 0, 0, 0}},
	}

	for _, c := range cases {
		out := threeSum(c.nums)
		fmt.Printf("Input: nums = %s \nOutput: %s \n\n", toJsonStr(c.nums), toJsonStr(out))
	}
}

func toJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// leetcode start

func threeSum(nums []int) [][]int {
	var out = [][]int{}
	if len(nums) < 3 {
		return out
	}

	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}

			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				out = append(out, []int{nums[i], nums[l], nums[r]})
				l++
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return out
}

func partition(list []int, i, j int) int {
	pivot := list[i]
	for i < j {
		if i < j && list[j] >= pivot {
			j--
		}
		list[i] = list[j]

		if i < j && list[i] <= pivot {
			i++
		}
		list[j] = list[i]
	}
	list[i] = pivot
	return i
}

func quickSort(list []int, i, j int) {
	if j > i {
		pivotIndex := partition(list, i, j)

		quickSort(list, i, pivotIndex-1)
		quickSort(list, pivotIndex+1, j)
	}
}
