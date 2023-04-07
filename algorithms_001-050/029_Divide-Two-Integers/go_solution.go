package main

import (
	"fmt"
	"math"
)

func main() {
	cases := []struct {
		dividend int
		divisor  int
	}{
		{dividend: 10, divisor: 3},
		{dividend: 7, divisor: -3},
		{dividend: math.MinInt32, divisor: -1},
		{dividend: 1, divisor: 1},
	}

	for _, c := range cases {
		fmt.Printf("Input: dividend = %d, divisor = %d \nOutput: %d \n\n", c.dividend, c.divisor, divide(c.dividend, c.divisor))
	}
}

// leetcode start

func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0 {
		return 0
	}
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}

	var symbol = 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		symbol = -1
	}

	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	out := 0
	total := dividend
	for total <= divisor {
		xPart := divisor
		xAddon := 1
		for (xPart<<1 >= total) && (xPart<<1 < xPart) {
			xPart <<= 1
			xAddon <<= 1
		}
		total -= xPart
		out += xAddon
	}

	return out * symbol
}
