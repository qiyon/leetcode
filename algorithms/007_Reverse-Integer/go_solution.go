package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d %d\n\n", math.MaxInt32, math.MinInt32)

	cases := []struct {
		x      int
		should int
	}{
		{x: -123, should: -321},
		{x: 120, should: 21},
		{x: 2147483647, should: 0},
		{x: 2147483641, should: 0},
		{x: -2147483648, should: 0},
		{x: 1000000008, should: 0},
	}

	for _, c := range cases {
		out := reverse(c.x)
		fmt.Printf("Input: x = %d \nOutput: %d \nShould: %d \n\n", c.x, out, c.should)
	}
}

// leetcode start

func reverse(x int) int {
	if x == math.MinInt32 {
		return 0
	}

	var (
		maxNums     = []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
		symbol  int = 1
		outNums []int
	)

	if x < 0 {
		symbol = -1
		x = -x
	}

	for x > 0 {
		item := x % 10
		outNums = append(outNums, item)
		x = (x - item) / 10
	}

	if len(outNums) >= len(maxNums) {
		for i, v := range outNums {
			if v < maxNums[i] {
				break
			}
			if v > maxNums[i] {
				return 0
			}
		}
	}

	outVal := 0
	for _, item := range outNums {
		outVal = outVal*10 + item
	}
	return symbol * outVal
}
