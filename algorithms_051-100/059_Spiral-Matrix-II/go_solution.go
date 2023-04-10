package main

import "fmt"

func main() {
	cases := []int{3, 1, 10}
	for _, c := range cases {
		out := generateMatrix(c)
		fmt.Printf("Input: n = %d\nOutput: \n", c)
		for _, row := range out {
			fmt.Printf("%v\n", row)
		}
		fmt.Println()
	}
}

// leetcode start

func generateMatrix(n int) [][]int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	l, r, top, bottom := 0, n-1, 0, n-1
	i, j := 0, 0
	iAdd, jAdd := 0, 1
	for v := 1; v <= n*n; v++ {
		grid[i][j] = v
		i += iAdd
		j += jAdd

		if jAdd == 1 && j == r {
			iAdd = 1
			jAdd = 0
			top++
		} else if iAdd == 1 && i == bottom {
			iAdd = 0
			jAdd = -1
			r--
		} else if jAdd == -1 && j == l {
			iAdd = -1
			jAdd = 0
			bottom--
		} else if iAdd == -1 && i == top {
			iAdd = 0
			jAdd = 1
			l++
		}
	}

	return grid
}
