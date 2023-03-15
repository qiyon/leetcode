package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		height []int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		{height: []int{1, 1}},
	}

	for _, c := range cases {
		out := maxArea(c.height)
		fmt.Printf("Input: height = %s\nOutput: %d\n\n", toJsonStr(c.height), out)
	}
}

func toJsonStr(v interface{}) string {
	j, _ := json.Marshal(v)
	return string(j)
}

// leetcode start

func maxArea(height []int) int {
	var (
		out    = 0
		lastIh = 0
		lastJh = 0
	)

	i, j := 0, len(height)-1
	for i < j {
		lastIh, lastJh = height[i], height[j]
		minH := lastIh
		if height[j] < minH {
			minH = lastJh
		}
		if area := minH * (j - i); area > out {
			out = area
		}

		// find next area, which may greater than last
		// min height's side && bigger than min height
		if lastIh < lastJh {
			for ; i < j && height[i] <= lastIh; i++ {
			}
		} else {
			for ; i < j && height[j] <= lastJh; j-- {
			}
		}
	}

	return out
}
