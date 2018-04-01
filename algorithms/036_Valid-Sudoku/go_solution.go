package main

import (
	"fmt"
)

//leetcode begin

func isValidSudoku(board [][]byte) bool {
	var rowHash = [9]map[byte]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	var colHash = [9]map[byte]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	var blockHash = [9]map[byte]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	for i, row := range board {
		for j, one_num := range row {
			if one_num == byte('.') {
				continue
			}
			_, ok := rowHash[i][one_num]
			if ok {
				return false
			} else {
				rowHash[i][one_num] = 1
			}
			_, ok2 := colHash[j][one_num]
			if ok2 {
				return false
			} else {
				colHash[j][one_num] = 1
			}
			block := (i/3)*3 + (j / 3)
			_, ok3 := blockHash[block][one_num]
			if ok3 {
				return false
			} else {
				blockHash[block][one_num] = 1
			}

		}
	}
	return true
}

//leectcode end

func main() {
	var soduku = [][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}
	fmt.Printf("return %v \n", isValidSudoku(soduku))
}
