package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d %d\n\n", math.MaxInt32, math.MinInt32)

	cases := []struct {
		s string
	}{
		{s: "42"},
		{s: "   -42"},
		{s: "4193 with words"},
		{s: "2147483648"},
		{s: "21474836460"},
		{s: "-2147483649"},
		{s: "00000-42a1234"},
	}

	for _, c := range cases {
		out := myAtoi(c.s)
		fmt.Printf("Input: s = \"%s\" \nOutput: %d \n\n", c.s, out)
	}
}

// leetcode start

func myAtoi(s string) int {
	var (
		symbol int
		nums   []int
	)

	for _, c := range s {
		if len(nums) == 0 && symbol == 0 {
			switch c {
			case ' ':
				continue
			case '+', '0':
				symbol = 1
				continue
			case '-':
				symbol = -1
				continue
			}
		}

		c2n := int(c - '0')
		if !(c2n >= 0 && c2n <= 9) {
			break
		}
		if c2n == 0 && len(nums) == 0 {
			continue
		}
		nums = append(nums, c2n)
	}

	if symbol == 0 {
		symbol = 1
	}

	maxNums := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	symbolMax := []int{-2147483648, 0, 2147483647}
	if len(nums) > len(maxNums) {
		return symbolMax[symbol+1]
	}
	if len(nums) == len(maxNums) {
		for i, v := range nums {
			if v < maxNums[i] {
				break
			}
			if v > maxNums[i] {
				return symbolMax[symbol+1]
			}
		}
	}

	out := 0
	for _, v := range nums {
		out = out*10 + v
	}
	return symbol * out
}
