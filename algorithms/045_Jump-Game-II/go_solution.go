package main

import "fmt"

//leetcode begin

func jump(nums []int) int {
	theLength := len(nums)
	if theLength <= 1 {
		return 0
	}
	mins := make([]int, theLength)
	mins[0] = 0
	hasRecordsIndex := 0
	for i := 0; i < theLength; i++ {
		if i > hasRecordsIndex {
			continue
		}
		steps := nums[i]
		yetStep := mins[i]
		for k := i + 1; k <= (i+steps) && k < theLength; k++ {
			if k > hasRecordsIndex {
				//fmt.Printf("%v %v %v \n", i, k, hasRecordsIndex)
				mins[k] = yetStep + 1
				hasRecordsIndex++
			} else {
				mins[k] = min(mins[k], yetStep+1)
			}
		}
	}
	return mins[theLength-1]
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

//leetcode end

func main() {
	tests := [][]int{
		{10, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 1},
		{0},
		{3, 2, 3, 4, 2, 5, 1},
	}
	for _, row := range tests {
		fmt.Printf("jump of: %v \nresult:%v \n", row, jump(row))
	}
}
