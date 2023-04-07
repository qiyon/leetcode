package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		nums1 []int
		nums2 []int
	}{
		{nums1: []int{}, nums2: []int{}},
		{nums1: []int{1}, nums2: []int{}},
		{nums1: []int{}, nums2: []int{3}},
		{nums1: []int{1, 3}, nums2: []int{2}},
		{nums1: []int{1, 2}, nums2: []int{3, 4}},
		{nums1: []int{1, 2, 3, 6, 7, 8}, nums2: []int{4, 5, 9, 10}},
	}

	for _, c := range cases {
		out := findMedianSortedArrays(c.nums1, c.nums2)
		fmt.Printf("Input: nums1 = %s, nums2 = %s\nOutput: %f\n\n", toJsonStr(c.nums1), toJsonStr(c.nums2), out)
	}
}

func toJsonStr(v interface{}) string {
	j, _ := json.Marshal(v)
	return string(j)
}

// leetcode start

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		length1 = len(nums1)
		length2 = len(nums2)
		l1      = 0
		l2      = 0
		r1      = length1 - 1
		r2      = length2 - 1
		vl, vr  int
	)

	if length1 == 0 && length2 == 0 {
		return 0
	}

	for l1 <= r1 || l2 <= r2 {
		if l1 <= r1 {
			if l2 <= r2 {
				if nums1[l1] < nums2[l2] {
					vl = nums1[l1]
					l1++
				} else {
					vl = nums2[l2]
					l2++
				}
				if nums1[r1] > nums2[r2] {
					vr = nums1[r1]
					r1--
				} else {
					vr = nums2[r2]
					r2--
				}
			} else {
				vl = nums1[l1]
				vr = nums1[r1]
				l1++
				r1--
			}
		} else {
			vl = nums2[l2]
			vr = nums2[r2]
			l2++
			r2--
		}
		// fmt.Printf("vl %d, vr %d \n", vl, vr)
	}

	return (float64(vl) + float64(vr)) / 2
}
