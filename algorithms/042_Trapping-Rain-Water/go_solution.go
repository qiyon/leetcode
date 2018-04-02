package main

import (
	"fmt"
)

//leetcode begin

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
		var maxIndex int
		var secondIndex int
		if f.height[l] > f.height[l+1] {
			maxIndex = l
			secondIndex = l + 1
		} else {
			maxIndex = l + 1
			secondIndex = l
		}
		for i := l + 2; i <= r; i++ {
			if f.height[i] > f.height[secondIndex] {
				if f.height[i] > f.height[maxIndex] {
					secondIndex = maxIndex
					maxIndex = i
				} else {
					secondIndex = i
				}
				//fmt.Printf("maxIndex: %v, secondIndex: %v \n", maxIndex, secondIndex)
			}
		}
		waterHeight := f.height[secondIndex]
		var nextL int
		var nextR int
		if maxIndex < secondIndex {
			nextL = maxIndex
			nextR = secondIndex
		} else {
			nextL = secondIndex
			nextR = maxIndex
		}
		//fmt.Printf("nextL: %v, nextR: %v \n", nextL, nextR)
		for j := nextL; j <= nextR; j++ {
			if f.height[j] < waterHeight {
				f.cnt += (waterHeight - f.height[j])
			}
		}
		f.calc(l, nextL)
		f.calc(nextR, r)
	}
}

//leetcode end

func main() {
	fmt.Printf("return %v \n", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("return %v \n", trap([]int{}))
	fmt.Printf("return %v \n", trap([]int{100, 2, 3, 4, 5, 6, 6, 5, 4, 2, 3, 2, 1}))
}
