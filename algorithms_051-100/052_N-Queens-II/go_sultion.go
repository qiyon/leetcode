package main

import "fmt"

//leetcode begin

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

//leetcode end

func main() {
	fmt.Printf("total N-Queens is: %v\n", totalNQueens(1))
	fmt.Printf("total N-Queens is: %v\n", totalNQueens(3))
	fmt.Printf("total N-Queens is: %v\n", totalNQueens(4))
	fmt.Printf("total N-Queens is: %v\n", totalNQueens(6))
	fmt.Printf("total N-Queens is: %v\n", totalNQueens(8))
	fmt.Printf("total N-Queens is: %v\n", totalNQueens(10))
}
