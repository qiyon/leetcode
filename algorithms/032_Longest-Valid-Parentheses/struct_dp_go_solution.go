package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		s string
	}{
		{s: ""},
		{s: "()"},
		{s: ")("},
		{s: ")(()())()"},
		{s: "()))"},
	}

	for _, c := range cases {
		out := longestValidParentheses(c.s)
		fmt.Printf("Input: s = \"%s\" \nOutput: %d \n\n", c.s, out)
	}
}

// leetcode start

type DpRecord struct {
	srcStr string
	srcLen int
	dpMap  []int
	max    int
}

func (s *DpRecord) init() {
	s.dpMap[0] = 0
}

func (s *DpRecord) runDp() {
	for index := 1; index < s.srcLen; index++ {
		preLen := s.dpMap[index-1]
		leftIndex := index - 1 - preLen
		if s.isParenthesesMatch(leftIndex, index) {
			s.dpMap[index] = s.dpMap[index-1] + 2
			//for scenes: ()(()) ==> 8
			if leftIndex-1 >= 0 {
				s.dpMap[index] += s.dpMap[leftIndex-1]
			}
		} else {
			s.dpMap[index] = 0
		}
		s.updateMax(s.dpMap[index])
	}
}

func (s *DpRecord) updateMax(maxVal int) {
	if s.max < maxVal {
		s.max = maxVal
	}
}

func (s *DpRecord) isParenthesesMatch(leftIndex int, rightIndex int) bool {
	return leftIndex >= 0 && s.srcStr[leftIndex] == '(' && s.srcStr[rightIndex] == ')'
}

func longestValidParentheses(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return 0
	}

	record := &DpRecord{
		srcStr: s,
		srcLen: sLen,
		dpMap:  make([]int, sLen),
		max:    0,
	}
	record.init()
	record.runDp()
	return record.max
}
