package main

import "fmt"

func main() {
	cases := []struct {
		x int
	}{
		{x: 121},
		{x: -121},
		{x: 10},
	}

	for _, c := range cases {
		out := isPalindrome(c.x)
		fmt.Printf("Input: s = %d \nOutput: %v \n\n", c.x, out)
	}
}

// leetcode start

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var nums []int
	for x > 0 {
		item := x % 10
		x = (x - item) / 10
		nums = append(nums, item)
	}

	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}

	return true
}
