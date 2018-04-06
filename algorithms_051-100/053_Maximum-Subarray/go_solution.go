package main

import "fmt"

//leetcode begin

func maxSubArray(nums []int) int {
	theLength := len(nums)
	if theLength == 0 {
		return 0
	}
	dp := make([]int, theLength)
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < theLength; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

//leetcode end

func main() {
	fmt.Printf("%v\n", maxSubArray([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
