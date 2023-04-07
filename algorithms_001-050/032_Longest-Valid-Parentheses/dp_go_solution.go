package main

import "fmt"

//leetcode begin
func longestValidParentheses_dp(s string) int {
	s_len := len(s)
	if s_len <= 1 {
		return 0
	}
	max := 0
	dp := make([]int, s_len)
	dp[0] = 0
	for i := 1; i < s_len; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			pre_len := dp[i-1]
			match_idx := i - 1 - pre_len
			if match_idx >= 0 && s[match_idx] == '(' {
				dp[i] = dp[i-1] + 2
				if match_idx-1 >= 0 {
					dp[i] += dp[match_idx-1]
				}
			} else {
				dp[i] = 0
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

//leetcode end
func main() {
	test_strs := []string{"()", ")(", ")(()())()", "()))"}
	for _, str := range test_strs {
		fmt.Printf("%v \n", str)
		max := longestValidParentheses_dp(str)
		fmt.Printf("%v \n", max)
	}
}
