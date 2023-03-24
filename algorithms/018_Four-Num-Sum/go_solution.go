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
		{nums: []int{1, 0, -1, 0, -2, 2}, target: 0},
		{nums: []int{2, 2, 2, 2, 2}, target: 8},
	}
	for _, c := range cases {
		out := fourSum(c.nums, c.target)
		fmt.Printf("Input: nums = %s, target = %d\nOutput: %s\n\n", toJsonStr(c.nums), c.target, toJsonStr(out))
	}
}

func toJsonStr(v interface{}) string {
	out, _ := json.Marshal(v)
	return string(out)
}

// leetcode start

func fourSum(nums []int, target int) [][]int {
	out := [][]int{}
	if len(nums) < 4 {
		return out
	}

	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					out = append(out, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}

		}
	}

	return out
}

func partition(list []int, i, j int) int {
	pivot := list[i]
	for i < j {
		for i < j && list[j] >= pivot {
			j--
		}
		list[i] = list[j]

		for i < j && list[i] <= pivot {
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
