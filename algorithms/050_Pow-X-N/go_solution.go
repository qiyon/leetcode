package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		x float64
		n int
	}{
		{x: 2.00000, n: 10},
		{x: 2.10000, n: 3},
		{x: 2.00000, n: -2},
		{x: 0, n: 0},
		{x: 100, n: 0},
		{x: 10, n: -2},
	}
	for _, c := range cases {
		out := myPow(c.x, c.n)
		fmt.Printf("Input: x = %.5f, n = %v\nOutput: %.5f\n\n", c.x, c.n, out)
	}
}

// leetcode start

func myPow(x float64, n int) float64 {
	if x == 0.0 {
		return 0.0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1.0 / x
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return x * myPow(x*x, n/2)
	}
}
