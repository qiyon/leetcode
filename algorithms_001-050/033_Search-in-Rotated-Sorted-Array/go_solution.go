package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
	for _, target := range nums {
		fmt.Printf("searched index: %v \n", search(nums, target))
	}
}

// leetcode start

type BinarySearch struct {
	nums       []int
	target     int
	left       int
	right      int
	mid        int
	leftValue  int
	midValue   int
	rightValue int
}

func (s *BinarySearch) search() {
	for {
		if s.isEnd() {
			break
		}
		s.calcuMidAndValue()
		if s.isAsc() {
			if s.isTargetInLeftMid() {
				s.right = s.mid
			} else {
				s.left = s.mid
			}
		} else {
			if s.isMidInLeftAsc() {
				if s.isTargetInLeftMid() {
					s.right = s.mid
				} else {
					s.left = s.mid
				}
			} else {
				if s.isTargetInMidRight() {
					s.left = s.mid
				} else {
					s.right = s.mid
				}
			}
		}
	}
}

func (s *BinarySearch) calcuMidAndValue() {
	s.mid = (s.left + s.right) / 2
	s.leftValue = s.nums[s.left]
	s.rightValue = s.nums[s.right]
	s.midValue = s.nums[s.mid]
}

func (s *BinarySearch) isAsc() bool {
	return s.leftValue < s.rightValue
}

func (s *BinarySearch) isMidInLeftAsc() bool {
	return s.midValue > s.leftValue
}

func (s *BinarySearch) isTargetInLeftMid() bool {
	return s.leftValue < s.target && s.target < s.midValue
}

func (s *BinarySearch) isTargetInMidRight() bool {
	return s.midValue < s.target && s.target < s.rightValue
}

func (s *BinarySearch) isEnd() bool {
	if s.left >= s.right-1 || s.nums[s.left] == s.target || s.nums[s.right] == s.target {
		return true
	} else {
		return false
	}
}

func (s BinarySearch) targetIndex() int {
	if s.nums[s.left] == s.target {
		return s.left
	}
	if s.nums[s.right] == s.target {
		return s.right
	}
	return -1
}
func search(nums []int, target int) int {
	cnt := len(nums)
	if cnt == 0 {
		return -1
	}
	binarySearch := BinarySearch{nums: nums, target: target, left: 0, right: cnt - 1}
	binarySearch.search()
	return binarySearch.targetIndex()
}
