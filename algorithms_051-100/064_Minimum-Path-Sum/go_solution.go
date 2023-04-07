package main

import "fmt"

func main() {
	cases := []struct {
		grid [][]int
	}{
		{grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}},
		{grid: [][]int{{1, 2, 3}, {4, 5, 6}}},
	}

	for _, c := range cases {
		out := minPathSum(c.grid)
		fmt.Printf("Input: grid = %v\nOutput: %d\n\n", c.grid, out)
	}
}

// leetcode start

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range grid {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
