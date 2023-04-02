package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		height []int
	}{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		{height: []int{4, 2, 0, 3, 2, 5}},
		{height: []int{}},
		{height: []int{100, 2, 3, 4, 5, 6, 6, 5, 4, 2, 3, 2, 1}},
	}

	for _, c := range cases {
		out := trap(c.height)
		fmt.Printf("Input: height = %v\nOutput: %d\n\n", c.height, out)
	}
}

// leetcode start

func trap(height []int) int {
	f := new(finder)
	f.height = height
	f.cnt = 0
	f.theLength = len(height)
	f.calc(0, f.theLength-1)
	return f.cnt
}

type finder struct {
	height    []int
	cnt       int
	theLength int
}

func (f *finder) calc(l int, r int) {
	if l < r-1 {
		var firstHIndex int
		var secondHIndex int
		if f.height[l] > f.height[l+1] {
			firstHIndex = l
			secondHIndex = l + 1
		} else {
			firstHIndex = l + 1
			secondHIndex = l
		}
		for i := l + 2; i <= r; i++ {
			if f.height[i] > f.height[secondHIndex] {
				if f.height[i] > f.height[firstHIndex] {
					secondHIndex = firstHIndex
					firstHIndex = i
				} else {
					secondHIndex = i
				}
			}
		}

		waterHeight := f.height[secondHIndex]
		var nextL int
		var nextR int
		if firstHIndex < secondHIndex {
			nextL = firstHIndex
			nextR = secondHIndex
		} else {
			nextL = secondHIndex
			nextR = firstHIndex
		}

		for j := nextL; j <= nextR; j++ {
			if f.height[j] < waterHeight {
				f.cnt += waterHeight - f.height[j]
			}
		}
		f.calc(l, nextL)
		f.calc(nextR, r)
	}
}
