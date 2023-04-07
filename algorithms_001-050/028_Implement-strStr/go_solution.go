package main

import "fmt"

func main() {
	cases := []struct {
		haystack string
		needle   string
	}{
		{"sadbutsad", "sad"},
		{"leetcode", "leeto"},
	}

	for _, c := range cases {
		fmt.Printf("Input: haystack = %s, needle = %s \nOutput: %v \n\n", c.haystack, c.needle, strStr(c.haystack, c.needle))
	}
}

// leetcode start

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
