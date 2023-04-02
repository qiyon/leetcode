package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		src [][]int
	}{
		{
			src: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
		},
		{
			src: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
		},
	}
	for _, c := range cases {
		fmt.Printf("Input:\n")
		for _, row := range c.src {
			fmt.Printf("%v\n", row)
		}

		rotate(c.src)

		fmt.Printf("Output:\n")
		for _, row := range c.src {
			fmt.Printf("%v\n", row)
		}
		fmt.Println()
	}
}

// leetcode start

func rotate(matrix [][]int) {
	l := new(loop)
	l.src = matrix
	l.theLength = len(matrix)
	l.loopRun(0, 0, l.theLength-2)
}

type loop struct {
	src       [][]int
	theLength int
}

func (l *loop) loopRun(row int, start int, end int) {
	if start > end {
		return
	}
	var tmpVal int
	for i := start; i <= end; i++ {
		x := row
		y := i
		tmpVal = l.src[x][y]
		for loop := 0; loop < 4; loop++ {
			//fmt.Printf("%v %v\n", x, y)
			x, y = y, l.theLength-1-x
			tmpVal, l.src[x][y] = l.src[x][y], tmpVal
		}
	}
	l.loopRun(row+1, start+1, end-1)
}
