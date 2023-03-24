package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		n int
	}{
		{n: 1},
		{n: 2},
		{n: 3},
		{n: 4},
	}

	for _, c := range cases {
		out := generateParenthesis(c.n)
		fmt.Printf("Input: n = %d\nOutput: %v\n\n", c.n, out)
	}
}

// leetcode start

func generateParenthesis(n int) []string {
	var (
		strStack     []string
		leftCntStack []int
		out          = []string{}
	)

	strStack = []string{"("}
	leftCntStack = []int{1}
	for len(strStack) > 0 {
		popStr := strStack[len(strStack)-1]
		strStack = strStack[:len(strStack)-1]
		popLeftCnt := leftCntStack[len(leftCntStack)-1]
		leftCntStack = leftCntStack[:len(leftCntStack)-1]

		if popLeftCnt < n {
			strStack = append(strStack, popStr+"(")
			leftCntStack = append(leftCntStack, popLeftCnt+1)
		}
		if popLeftCnt > (len(popStr) - popLeftCnt) {
			if len(popStr) == 2*n-1 {
				out = append(out, popStr+")")
			} else {
				strStack = append(strStack, popStr+")")
				leftCntStack = append(leftCntStack, popLeftCnt)
			}
		}
	}
	return out
}
