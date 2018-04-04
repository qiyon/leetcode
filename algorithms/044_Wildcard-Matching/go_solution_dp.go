package main

import "fmt"

//leetcoe begin

func isMatch(s string, p string) bool {
	m := new(matcher)
	m.s = []byte(s)
	m.sLength = len(m.s)
	m.p = make([]byte, len(p))
	m.pLength = 0
	for _, val := range []byte(p) {
		//remove char star repeat
		if m.pLength > 0 && val == byte('*') && m.p[m.pLength-1] == byte('*') {
			continue
		}
		m.p[m.pLength] = val
		m.pLength++
	}
	m.dp = make([][]bool, m.sLength+1)
	for i := 0; i <= m.sLength; i++ {
		m.dp[i] = make([]bool, m.pLength+1)
	}
	m.toMatch()
	return m.dp[m.sLength][m.pLength]
}

type matcher struct {
	s       []byte
	p       []byte
	sLength int
	pLength int
	dp      [][]bool
}

func (m *matcher) toMatch() {
	m.dp[0][0] = true
	for i := 0; i <= m.sLength; i++ {
		for j := 1; j <= m.pLength; j++ {
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
		if m.getDpStatus(i-1, j-1) && ( m.s[i-1] == m.p[j-1] || m.p[j-1] == byte('?') ) {
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

//leetcode end

func main() {
	in := [][]string{
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
	for _, row := range in {
		fmt.Printf("match '%v' with '%v' result: %v \n", row[0], row[1], isMatch(row[0], row[1]))
	}
}
