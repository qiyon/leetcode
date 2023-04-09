package main

import "fmt"

func main() {
	cases := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
		"   ",
		"abc   ",
	}

	for _, c := range cases {
		out := lengthOfLastWord(c)
		fmt.Printf("Input: s = \"%s\"\nOutput: %d\n\n", c, out)
	}
}

// leetcode start

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	startIdx := len(s) - 1
	hasTrim := false
	trimCnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		startIdx = i
		if hasTrim == false && s[i] == ' ' {
			trimCnt++
			continue
		}

		hasTrim = true
		if hasTrim && s[i] == ' ' {
			startIdx = i + 1
			break
		}
	}

	return len(s) - startIdx - trimCnt
}
