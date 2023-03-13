package main

import "fmt"

func main() {
	cases := []string{
		"babad",
		"cbbd",
		"aaaaa",
		"aaaaaa",
		"aaaaaaa",
		"aaaaaaaa",
		"aaabaaaa",
	}

	for _, c := range cases {
		out := longestPalindrome(c)
		fmt.Printf("Input: s = %s\nOutput: %s\n\n", c, out)
	}
}

// leetcode start. Approach 4: Expand Around Center

func longestPalindrome(s string) string {
	var (
		longest int
		longEnd int
	)

	if s == "" {
		return ""
	}

	for i, _ := range s {
		l, r := expandCenter(s, i, i)
		if r-l-1 > longest {
			longest = r - l - 1
			longEnd = r - 1
		}

		l, r = expandCenter(s, i, i+1)
		if r-l-1 > longest {
			longest = r - l - 1
			longEnd = r - 1
		}
	}

	return s[longEnd-longest+1 : longEnd+1]
}

func expandCenter(s string, lIn, rIn int) (int, int) {
	var (
		lOut = lIn
		rOut = rIn
	)

	for lOut >= 0 && rOut < len(s) && s[lOut] == s[rOut] {
		lOut--
		rOut++
	}

	return lOut, rOut
}

// leetcode start. Approach 3: Dynamic Programming

func longestPalindrome_dp(s string) string {
	var (
		longest int
		longEnd int
	)

	if s == "" {
		return ""
	}

	longest = 1
	longEnd = 0

	// define dp[i][j]bool: s[i:j+1] is a palindrome
	// init:
	// 	 dp[i][i] is true
	// 	 s[i]=s[i+1] >= dp[i][i+1] is true
	dp := make([][]bool, len(s))
	for i, _ := range s {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		if i > 0 && s[i] == s[i-1] {
			dp[i-1][i] = true
			if longest < 2 {
				longest = 2
				longEnd = i
			}
		}
	}

	// dp[i][j] = dp[i+1][j-1] && s[i]=s[j]
	// loop:
	//   i=o-r,j=o+r
	//   i=o-r,j=o+1+r
	for o, _ := range s {
		for r := 1; o-r >= 0 && o+r < len(s); r++ {
			if dp[o-r+1][o+r-1] && s[o-r] == s[o+r] {
				dp[o-r][o+r] = true
				if 2*r+1 > longest {
					longest = 2*r + 1
					longEnd = o + r
				}
			}
			if o+1+r < len(s) {
				if dp[o-r+1][o+1+r-1] && s[o-r] == s[o+1+r] {
					dp[o-r][o+1+r] = true
					if 2*r+2 > longest {
						longest = 2*r + 2
						longEnd = o + 1 + r
					}
				}
			}
		}
	}

	return s[longEnd-longest+1 : longEnd+1]
}
