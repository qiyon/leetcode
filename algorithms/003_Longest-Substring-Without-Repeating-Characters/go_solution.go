package main

import "fmt"

func main() {
	cases := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		" ",
	}

	for _, c := range cases {
		out := lengthOfLongestSubstring(c)
		fmt.Printf("Input: s = %s \nOutput: %d\n\n", c, out)
	}
}

// leetcode start

func lengthOfLongestSubstring(s string) int {
	var (
		charMap  = map[uint8]int{}
		indexMax = make([]int, len(s))
	)

	if s == "" {
		return 0
	}

	charMap[s[0]] = 0
	indexMax[0] = 1
	max := 1
	for i := 1; i < len(s); i++ {
		v := s[i]
		length := indexMax[i-1] + 1
		if preIndex, ok := charMap[v]; ok {
			if i-preIndex < length {
				length = i - preIndex
			}
		}

		if length > max {
			max = length
		}

		charMap[v] = i
		indexMax[i] = length
	}
	return max
}
