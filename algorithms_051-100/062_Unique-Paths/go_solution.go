package main

import "fmt"

func main() {
	cases := []struct {
		m int
		n int
	}{
		{m: 3, n: 7},
		{m: 3, n: 2},
		{m: 3, n: 3},
		{m: 16, n: 16},
	}

	for _, c := range cases {
		out := uniquePaths(c.m, c.n)
		fmt.Printf("Input: m = %d, n = %d\nOutput: %d\n\n", c.m, c.n, out)
	}
}

// leetcode start

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	cn := int64(m + n - 2)
	cr := int64(m - 1)
	if n < m {
		cr = int64(n - 1)
	}

	var x, y int64
	x, y = 1, 1
	for cr >= 1 {
		x *= cn
		y *= cr
		for x%2 == 0 && y%2 == 0 {
			x /= 2
			y /= 2
		}
		cn--
		cr--
	}

	return int(x / y)
}
