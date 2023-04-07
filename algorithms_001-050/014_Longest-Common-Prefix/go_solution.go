package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		strs []string
	}{
		{strs: []string{"flower", "flow", "flight"}},
		{strs: []string{"dog", "racecar", "car"}},
	}

	for _, c := range cases {
		out := longestCommonPrefix(c.strs)
		fmt.Printf("Input: strs = %v \nOutput: %s \n\n", c.strs, out)
	}
}

// leetcode start

func longestCommonPrefix(strs []string) string {
	var (
		out string
	)

	if len(strs) == 0 {
		return ""
	}

	for i, _ := range strs[0] {
		for j, _ := range strs {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return out
			}
		}
		out += string(strs[0][i])
	}

	return out
}
