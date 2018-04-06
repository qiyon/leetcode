package main

import "fmt"

//leetcode begin

func canJump(nums []int) bool {
	theLength := len(nums)
	if theLength <= 1 {
		return true
	}
	last := theLength - 1
	for i := theLength - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last <= 0
}

//leetcode end

func main() {
	fmt.Printf("%v\n", canJump([]int{2, 3, 1, 1, 4}))
	fmt.Printf("%v\n", canJump([]int{3, 2, 1, 0, 4}))
}
