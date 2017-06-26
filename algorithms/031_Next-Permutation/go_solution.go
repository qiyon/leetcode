package main

import "fmt"

//--leetcode begin
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
	exchange_idx := plode + 1
	for i := plode + 1; i < cnt; i++ {
		if nums[i] > hiil && nums[i] <= nums[exchange_idx] {
			exchange_idx = i
		}
	}
	//swap
	nums[plode], nums[exchange_idx] = nums[exchange_idx], nums[plode]
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

//--leetcode end
func main() {
	nums := []int{1, 2, 3, 4}
	for i := 0; i < 100; i++ {
		nextPermutation(nums)
		fmt.Printf("%v \n", nums)
	}
}
