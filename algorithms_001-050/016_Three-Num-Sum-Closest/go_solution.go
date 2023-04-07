package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
	}

	for _, c := range cases {
		out := threeSumClosest(c.nums, c.target)
		fmt.Printf("Input: nums = [%v], target = %d\nOutput: %d\nExpected: %d\n\n", c.nums, c.target, out, c.expected)
	}
}

// leetcode start

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	out := nums[0] + nums[1] + nums[2]

	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			}

			if abs(sum-target) < abs(out-target) {
				out = sum
			}
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return out
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
