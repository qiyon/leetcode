package main

import (
	"fmt"
)

func main() {
	cases := []int{1, 3, 4, 6, 8, 10}
	for _, c := range cases {
		out := totalNQueens(c)
		fmt.Printf("Input: n = %d\nOutput: %d\n\n", c, out)
	}
}

// leetcode start

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return 0
	}
	q := new(queens)
	q.n = n
	q.res = make([][]int, 0)
	q.cnt = 0
	s := new(stat)
	s.list = make([]int, n)
	s.theLength = 0
	q.run(s)
	return q.cnt
}

type queens struct {
	n   int
	res [][]int
	cnt int
}

func (q *queens) run(s *stat) {
	if s.theLength == q.n {
		q.cnt++
		return
	}
	for col := 0; col < q.n; col++ {
		if q.isValid(s, col) {
			s.addQueen(col)
			q.run(s)
			s.pop()
		}
	}
}

func (q *queens) isValid(s *stat, col int) bool {
	row := s.theLength
	for preRow := 0; preRow < s.theLength; preRow++ {
		if s.list[preRow] == col {
			return false
		}
		if s.list[preRow]-preRow == col-row {
			return false
		}
		if s.list[preRow]+preRow == col+row {
			return false
		}
	}
	return true
}

type stat struct {
	//for each row, the index of queen
	list      []int
	theLength int
}

func (s *stat) addQueen(col int) {
	s.list[s.theLength] = col
	s.theLength++
}

func (s *stat) pop() {
	s.theLength--
}
