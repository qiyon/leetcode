package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		matrix [][]int
	}{
		{matrix: [][]int{
			{1},
		}},
		{matrix: [][]int{
			{1, 2, 3},
		}},
		{matrix: [][]int{
			{1},
			{2},
			{3},
		}},
		{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}},
		{matrix: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}},
		{matrix: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		}},
		{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{10, 11, 12},
		}},
	}
	for _, c := range cases {
		out := spiralOrder(c.matrix)
		fmt.Printf("Input: matrix= %v\nOutput:%v\n\n", c.matrix, out)
	}
}

// leetcode start

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])
	cnt := m * n
	r := new(res)
	r.list = make([]int, cnt)
	r.index = 0
	row := 0
	col := 0
	minRow := 0
	minCol := 0
	maxRow := n - 1
	maxCol := m - 1
	colAdd := 1
	rowAdd := 0
	for r.index < cnt {
		//fmt.Printf("%v %v %v %v\n", row, col, rowAdd, colAdd)
		r.add(matrix[row][col])
		if colAdd == 1 && col == maxCol {
			colAdd = 0
			rowAdd = 1
		} else if rowAdd == 1 && row == maxRow {
			rowAdd = 0
			colAdd = -1
		} else if colAdd == -1 && col == minCol {
			colAdd = 0
			rowAdd = -1
		} else if rowAdd == -1 && row == minRow+1 {
			colAdd = 1
			rowAdd = 0
			minRow++
			minCol++
			maxRow--
			maxCol--
		}
		row += rowAdd
		col += colAdd
	}
	return r.list
}

type res struct {
	list  []int
	index int
}

func (r *res) add(val int) {
	r.list[r.index] = val
	r.index++
}
