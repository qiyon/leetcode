package main

import "fmt"

// Leetcode begin
type BinarySearch struct {
	nums   []int
	cnt    int
	target int
	left   int
	right  int
}

func (self *BinarySearch) search() {
	for {
		if self.isEnd() {
			break
		}
		mid := (self.left + self.right) / 2
		if self.leftVal() < self.rightVal() {
			if self.midVal(mid) > self.target {
				self.right = mid
			} else {
				self.left = mid
			}
		} else {
			if self.midVal(mid) > self.leftVal() {
				if self.leftVal() < self.target && self.target < self.midVal(mid) {
					self.right = mid
				} else {
					self.left = mid
				}
			} else {
				if self.midVal(mid) < self.target && self.target < self.rightVal() {
					self.left = mid
				} else {
					self.right = mid
				}
			}
		}
	}
}

func (self BinarySearch) leftVal() int {
	return self.nums[self.left]
}

func (self BinarySearch) rightVal() int {
	return self.nums[self.right]
}

func (self BinarySearch) midVal(mid int) int {
	return self.nums[mid%self.cnt]
}

func (self BinarySearch) isEnd() bool {
	if self.left >= self.right-1 || self.nums[self.left] == self.target || self.nums[self.right] == self.target {
		return true
	} else {
		return false
	}
}

func (self *BinarySearch) targetIndex() int {
	if self.nums[self.left] == self.target {
		return self.left
	}
	if self.nums[self.right] == self.target {
		return self.right
	}
	return -1
}

func search(nums []int, target int) int {
	cnt := len(nums)
	if cnt == 0 {
		return -1
	}
	bin_search := BinarySearch{nums: nums, cnt: cnt, target: target, left: 0, right: cnt - 1}
	bin_search.search()
	return bin_search.targetIndex()
}

// Leetcode end

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2, 3}
	for _, target := range arr {
		fmt.Printf("searched index: %v \n", search(arr, target))
	}
}
