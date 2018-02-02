package main

import "fmt"

// Leetcode begin
type BinarySearch struct {
	nums        []int
	target      int
	left        int
	right       int
	mid         int
	left_value  int
	mid_value   int
	right_value int
}

func (self *BinarySearch) search() {
	for {
		if self.isEnd() {
			break
		}
		self.calcuMidAndValue()
		if self.isAsc() {
			if self.isTargetInLeftMid() {
				self.right = self.mid
			} else {
				self.left = self.mid
			}
		} else {
			if self.isMidInLeftAsc() {
				if self.isTargetInLeftMid() {
					self.right = self.mid
				} else {
					self.left = self.mid
				}
			} else {
				if self.isTargetInMidRight() {
					self.left = self.mid
				} else {
					self.right = self.mid
				}
			}
		}
	}
}

func (self *BinarySearch) calcuMidAndValue() {
	self.mid = (self.left + self.right) / 2
	self.left_value = self.nums[self.left]
	self.right_value = self.nums[self.right]
	self.mid_value = self.nums[self.mid]
}

func (self BinarySearch) isAsc() bool {
	return self.left_value < self.right_value
}

func (self BinarySearch) isMidInLeftAsc() bool {
	return self.mid_value > self.left_value
}

func (self BinarySearch) isTargetInLeftMid() bool {
	return self.left_value < self.target && self.target < self.mid_value
}

func (self BinarySearch) isTargetInMidRight() bool {
	return self.mid_value < self.target && self.target < self.right_value
}

func (self BinarySearch) isEnd() bool {
	if self.left >= self.right-1 || self.nums[self.left] == self.target || self.nums[self.right] == self.target {
		return true
	} else {
		return false
	}
}

func (self BinarySearch) targetIndex() int {
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
	bin_search := BinarySearch{nums: nums, target: target, left: 0, right: cnt - 1}
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
