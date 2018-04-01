package main

import (
	"fmt"
)

//leetcode start

func firstMissingPositive(nums []int) int {
	theLength := len(nums)
	//fmt.Printf("%v \n", nums)
	for i := 0; i < theLength; i++ {
		for nums[i] > 0 && nums[i] <= theLength && nums[i] != i+1 {
			toPosition := nums[i] - 1
			if nums[toPosition] == toPosition+1 {
				break
			}
			nums[i], nums[toPosition] = nums[toPosition], nums[i]
			//fmt.Printf("%v \n", nums)
		}
	}
	for k, val := range nums {
		if val != k+1 {
			return k + 1
		}
	}
	return theLength + 1
}

//leetcode end

func main() {
	fmt.Printf("return: %v \n", firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Printf("return: %v \n", firstMissingPositive([]int{1, 2, 0}))
	fmt.Printf("return: %v \n", firstMissingPositive([]int{1, 2, 0, 2, 2, -2, 2, 2, 2, 4, 5, 6, 7, 8}))
	fmt.Printf("return: %v \n", firstMissingPositive([]int{1, 2, 3}))
}
