package main

import "fmt"

//leetcode begin
func longestValidParentheses(s string) int {
    s_len := len(s)
    if s_len <= 1 {
        return 0
    }
    max := 0
    dp := make([]int, s_len)
    dp[0] = 0
    for i := 1; i < s_len-1; i++ {
        if s[i] == '(' {
            dp[i] = 0
        } else {
            pre_len := dp[i - 1]
            if s[i - 1 - pre_len] == '(' {
                dp[i] = dp[i-1] + 2
                if i - 1 - pre_len - 1 >= 0 {
                    dp[i] += dp[i - 1 - pre_len - 1]
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
    str := ")()(()())("
    fmt.Printf("%v \n", str)
    max := longestValidParentheses(str)
    fmt.Printf("%v", max)
}
