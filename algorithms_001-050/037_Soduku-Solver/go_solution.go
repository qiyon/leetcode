package main

import (
	"fmt"
)

func main() {
	var soduku = [][]byte{
		[]byte("..9748..."),
		[]byte("7........"),
		[]byte(".2.1.9..."),
		[]byte("..7...24."),
		[]byte(".64.1.59."),
		[]byte(".98...3.."),
		[]byte("...8.3.2."),
		[]byte("........6"),
		[]byte("...2759.."),
	}
	// soduku = [][]byte{
	// 	[]byte("........."),
	// 	[]byte("........."),
	// 	[]byte("........."),
	// 	[]byte("........."),
	// 	[]byte("........."),
	// 	[]byte("........."),
	// 	[]byte("........."),
	// 	[]byte("........."),
	// 	[]byte("........."),
	// }
	solveSudoku(soduku)
	for _, row := range soduku {
		for _, val := range row {
			fmt.Printf(string(rune(val)))
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}

func (d dfs) debug(now_i int, now_j int, c byte) {
	fmt.Printf("now_i: %v, now_j:%v, byte:%v \n", now_i, now_j, c-byte('0'))
	fmt.Printf("i: %v, j:%v \n", d.i, d.j)
	for _, row := range d.b {
		for _, val := range row {
			fmt.Printf(string(rune(val)))
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("------------------\n")
}

// leetcode start

func solveSudoku(board [][]byte) {
	d := new(dfs)
	d.src = board
	d.init()
	d.search()
	for i, row := range d.b {
		for j, val := range row {
			board[i][j] = val
		}
	}
}

type dfs struct {
	src [][]byte
	b   [9][9]byte
	i   int
	j   int
}

func (d *dfs) init() {
	for i, row := range d.src {
		for j, val := range row {
			d.b[i][j] = val
		}
	}
	d.i = 0
	d.j = 0
}

func (d *dfs) search() {
	now_i := d.i
	now_j := d.j
	now_byte := d.b[d.i][d.j]
	if now_byte != '.' {
		if d.isEnd() {
			return
		}
		// d.debug(now_i, now_j, 'now_byte')
		d.next(now_i, now_j)
		d.search()
	} else {
		mat := d.foundMat(now_i, now_j)
		for _, right_byte := range []byte("123456789") {
			letfed_val := (1 << (right_byte - byte('0')))
			if (letfed_val & mat) == 0 {
				// d.debug(now_i, now_j, right_byte)
				if d.isEnd() {
					if now_i == 8 && now_j == 8 {
						d.b[now_i][now_j] = right_byte
					}
					return
				}
				d.b[now_i][now_j] = right_byte
				d.next(now_i, now_j)
				d.search()
			}
		}
	}
	if d.isEnd() {
		return
	}
	d.i = now_i
	d.j = now_j
	d.b[now_i][now_j] = now_byte
}

func (d *dfs) next(i int, j int) {
	if j == 8 {
		d.j = 0
		d.i = i + 1
	} else {
		d.j = j + 1
	}
}

func (d dfs) isEnd() bool {
	return d.i == 8 && d.j == 8
}

func (d dfs) foundMat(i int, j int) int {
	mat := 1
	for _, val := range d.b[i] {
		if val == byte('.') {
			continue
		}
		mat = mat | (1 << (val - byte('0')))
	}
	for _, row := range d.b {
		if row[j] == byte('.') {
			continue
		}
		mat = mat | (1 << (row[j] - byte('0')))
	}
	for _, b_i := range []int{0, 1, 2} {
		for _, b_j := range []int{0, 1, 2} {
			if d.b[(i/3)*3+b_i][(j/3)*3+b_j] == byte('.') {
				continue
			}
			mat = mat | (1 << (d.b[(i/3)*3+b_i][(j/3)*3+b_j] - byte('0')))
		}
	}
	return mat
}
