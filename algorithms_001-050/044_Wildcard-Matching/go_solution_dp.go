package main

import (
	"fmt"
)

func main() {
	cases := [][]string{
		{"0cvd", "??*"},
		{"abcd", "ab*d"},
		{"a", "*"},
		{"sdf", "sdf"},
		{"sdfs", "***?"},
		{"44", "??"},
		{"", "***"},
		{"axxbcd", "axxbcd*"},
		{"a", "?*"},
		{"c", "*?*"},
		{"abcd", "a*bdc"},
		{"axxbcd", "a*c"},
		{"axxbcd", "axxbcd*cd"},
		{
			"babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb",
			"b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a",
		},
	}
	for _, c := range cases {
		fmt.Printf("Input: s= '%v' p= '%v' \nOutput: %v \n\n", c[0], c[1], isMatch(c[0], c[1]))
	}
}

// leetcode start

func isMatch(s string, p string) bool {
	m := &matcher{
		s:    []byte(s),
		sLen: len(s),
		p:    make([]byte, len(p)),
		pLen: 0,
	}
	for _, val := range []byte(p) {
		//remove char star repeat
		if m.pLen > 0 && val == byte('*') && m.p[m.pLen-1] == byte('*') {
			continue
		}
		m.p[m.pLen] = val
		m.pLen++
	}

	m.dp = make([][]bool, m.sLen+1)
	for i := 0; i <= m.sLen; i++ {
		m.dp[i] = make([]bool, m.pLen+1)
	}

	m.toMatch()
	return m.dp[m.sLen][m.pLen]
}

type matcher struct {
	s    []byte
	p    []byte
	sLen int
	pLen int
	dp   [][]bool
}

func (m *matcher) toMatch() {
	m.dp[0][0] = true
	for i := 0; i <= m.sLen; i++ {
		for j := 1; j <= m.pLen; j++ {
			m.recordDpStatus(i, j, m.checkStatusByDp(i, j))
		}
	}
}

func (m *matcher) checkStatusByDp(i int, j int) bool {
	if m.getDpStatus(i, j-1) && m.p[j-1] == byte('*') {
		return true
	}
	if i > 0 {
		if m.p[j-1] == byte('*') && (m.getDpStatus(i-1, j-1) || m.getDpStatus(i-1, j)) {
			return true
		}
		if m.getDpStatus(i-1, j-1) && (m.s[i-1] == m.p[j-1] || m.p[j-1] == byte('?')) {
			return true
		}
	}
	return false
}

func (m *matcher) recordDpStatus(i int, j int, val bool) {
	m.dp[i][j] = val
}

func (m *matcher) getDpStatus(i int, j int) bool {
	return m.dp[i][j]
}
