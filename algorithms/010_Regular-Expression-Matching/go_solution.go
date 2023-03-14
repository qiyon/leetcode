package main

import "fmt"

func main() {
	cases := []struct {
		s string
		p string
	}{
		{s: "abc", p: "abc"},
		{s: "aa", p: "a"},
		{s: "aab", p: "c*a*b"},
		{s: "aa", p: "a*"},
		{s: "aaa", p: "ab*"},
		{s: "ab", p: ".*"},
		{s: "mississippi", p: "mis*is*p*."},
	}

	for _, c := range cases {
		out := isMatch(c.s, c.p)
		fmt.Printf("Input: s = \"%s\", p = \"%s\"\nOutput: %v \n\n", c.s, c.p, out)
	}
}

// leetcode start

func isMatch(s string, p string) bool {
	// dp[i][j] => s[:i] p[:j] is match
	var dp = make([][]bool, len(s)+1)

	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	// dp init, s[:0] p[:0] is match
	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		jChar := p[j-1]
		for i := 0; i <= len(s); i++ {
			// loop
			// dp[i][j]
			//   dp[i-1][j-1] && char match
			//   or * case: zero, one, repeat
			switch jChar {
			case '*':
				if j < 2 {
					continue
				}
				preChar := p[j-2]
				if preChar == '*' {
					continue
				}
				dp[i][j] = dp[i][j-2] || (i > 0 && (dp[i-1][j-2] || dp[i-1][j]) && (preChar == '.' || preChar == s[i-1]))
			case '.':
				dp[i][j] = i > 0 && dp[i-1][j-1]
			default:
				dp[i][j] = i > 0 && dp[i-1][j-1] && s[i-1] == jChar
			}
			// fmt.Printf("dp[%d][%d]=%v\n", i, j, dp[i][j])
		}
	}

	return dp[len(s)][len(p)]
}
