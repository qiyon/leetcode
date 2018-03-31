package main

import "fmt"

//leetcode begin
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	if end < 0 {
		return 0
	}
	if nums[0] > target {
		return 0
	}
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
    if nums[end] > target {
        return  start + 1
    }else{
        return end +1
    }
}
//leetcode end

func main() {
    arr := []int{1, 3, 5, 6, 7, 8, 10}
    for _, v := range arr {
        fmt.Printf("return %v \n", searchInsert(arr, v))
    }
    fmt.Printf("return %v \n", searchInsert(arr, 0))
    fmt.Printf("return %v \n", searchInsert(arr, 2))
    fmt.Printf("return %v \n", searchInsert(arr, 4))
    fmt.Printf("return %v \n", searchInsert(arr, 9))
    fmt.Printf("return %v \n", searchInsert(arr, 11))
    fmt.Printf("return %v \n", searchInsert([]int{}, 0))
    fmt.Printf("return %v \n", searchInsert([]int{1}, 0))
    fmt.Printf("return %v \n", searchInsert([]int{1}, 2))
}
