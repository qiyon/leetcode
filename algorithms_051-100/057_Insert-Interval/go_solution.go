package main

import "fmt"

func main() {
	cases := []struct {
		intervals   [][]int
		newInterval []int
	}{
		{intervals: [][]int{{1, 3}, {6, 9}}, newInterval: []int{2, 5}},
		{intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}},
		{intervals: [][]int{{1, 5}}, newInterval: []int{2, 7}},
	}

	for _, c := range cases {
		out := insert(c.intervals, c.newInterval)
		fmt.Printf("Input: intervals = %v, newInterval = %v\nOutput: %v\n\n", c.intervals, c.newInterval, out)
	}
}

// leetcode start

func insert(intervals [][]int, newInterval []int) [][]int {
	out := [][]int{}
	if len(intervals) == 0 {
		out = append(out, newInterval)
		return out
	}

	intervalFirst := intervals[0]
	if newInterval[0] < intervalFirst[0] {
		intervalFirst, newInterval = newInterval, intervalFirst
	}
	out = append(out, intervalFirst)
	inserted := false
	i := 1
	for i < len(intervals) || !inserted {
		addon := []int{0, 0}
		if !inserted && (i >= len(intervals) || newInterval[0] <= intervals[i][0]) {
			addon = newInterval
			inserted = true
		} else {
			addon = intervals[i]
			i++
		}

		endIdx := len(out) - 1
		if addon[1] <= out[endIdx][1] {
			continue
		}
		if addon[0] <= out[endIdx][1] {
			out[endIdx][1] = addon[1]
		} else {
			out = append(out, addon)
		}
	}

	return out
}
