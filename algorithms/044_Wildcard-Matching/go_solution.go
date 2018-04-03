package main

import "fmt"

//leetcoe begin

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	m := new(matcher)
	m.s = []byte(s)
	m.p = make([]byte, len(p))
	m.sMaxIndex = len(m.s) - 1
	m.pMaxIndex = -1
	for _, val := range []byte(p) {
		//remove char star repeat
		if m.pMaxIndex >= 0 && val == byte('*') && m.p[m.pMaxIndex] == byte('*') {
			continue
		}
		m.pMaxIndex++
		m.p[m.pMaxIndex] = val
	}
	if m.pMaxIndex == 0 && m.p[0] == byte('*') {
		return true
	}
	if m.sMaxIndex == -1 {
		return false
	}
	m.dp = make([][]bool, m.sMaxIndex+2)
	for i := 0; i <= m.sMaxIndex+1; i++ {
		m.dp[i] = make([]bool, m.pMaxIndex+2)
		for j := 0; j <= m.pMaxIndex+1; j++ {
			m.dp[i][j] = true
		}
	}
	return m.toMatch(0, 0)
}

type matcher struct {
	s         []byte
	p         []byte
	sMaxIndex int
	pMaxIndex int
	dp        [][]bool
}

func (m *matcher) toMatch(i int, j int) bool {
	//fmt.Printf("%v %v \n", i, j)
	if m.dp[i][j] == false {
		return false
	}
	if i > m.sMaxIndex {
		if j == m.pMaxIndex && m.p[j] == byte('*') {
			return true
		}
		m.recordDpFalse(i, j)
		return false
	}
	if m.p[j] == byte('*') {
		if j == m.pMaxIndex {
			return true
		}
		if m.toMatch(i, j+1) {
			return true
		}
		if m.toMatch(i+1, j) {
			return true
		}
		if m.toMatch(i+1, j+1) {
			return true
		}
	}
	if m.s[i] == m.p[j] || m.p[j] == byte('?') {
		if i == m.sMaxIndex && j == m.pMaxIndex {
			return true
		}
		if j < m.pMaxIndex && m.toMatch(i+1, j+1) {
			return true
		}
	}
	m.recordDpFalse(i, j)
	return false
}

func (m *matcher) recordDpFalse(i int, j int) {
	m.dp[i][j] = false
}

//leetcode end

func main() {
	in := [][]string{
		{"abcd", "ab*d"},
		{"a", "*"},
		{"0cvd", "??*"},
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
	for _, row := range in {
		fmt.Printf("match '%v' with '%v' result: %v \n", row[0], row[1], isMatch(row[0], row[1]))
	}
}
