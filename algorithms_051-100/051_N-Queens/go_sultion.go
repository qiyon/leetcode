package main

import (
	"fmt"
)

func main() {
	cases := []int{5}
	for _, c := range cases {
		out := solveNQueens(c)
		fmt.Printf("Input: n = %d\nOutput: \n", c)
		for _, r := range out {
			for _, row := range r {
				fmt.Printf("%v\n", row)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

// leetcode start

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n <= 3 {
		return [][]string{}
	}
	q := new(queens)
	q.n = n
	q.res = make([][]int, 0)
	s := new(stat)
	s.list = make([]int, n)
	s.theLength = 0
	q.run(s)
	return q.genResult()
}

type queens struct {
	n   int
	res [][]int
}

func (q *queens) run(s *stat) {
	if s.theLength == q.n {
		q.addResult(s)
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

func (q *queens) addResult(s *stat) {
	tmp := make([]int, q.n)
	copy(tmp, s.list)
	q.res = append(q.res, tmp)
}

func (q *queens) genResult() [][]string {
	resList := make([][]string, 0)
	for _, solution := range q.res {
		resItem := make([]string, q.n)
		for row, qPositon := range solution {
			resRow := make([]byte, q.n)
			for col := 0; col < q.n; col++ {
				if qPositon == col {
					resRow[col] = byte('Q')
				} else {
					resRow[col] = byte('.')
				}
			}
			resItem[row] = string(resRow)
		}
		resList = append(resList, resItem)
	}
	return resList
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
