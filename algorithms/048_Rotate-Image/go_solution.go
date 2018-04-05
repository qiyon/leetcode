package main

import "fmt"

//leetcode begin

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

//leetcode end

func main() {
	src := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(src)
	for _, row := range src {
		fmt.Printf("%v\n", row)
	}
	src2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(src2)
	for _, row := range src2 {
		fmt.Printf("%v\n", row)
	}
}
