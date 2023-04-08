package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		intervals [][]int
	}{
		{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
		{intervals: [][]int{{1, 4}, {4, 5}}},
		{intervals: [][]int{{1, 4}, {2, 3}}},
	}

	for _, c := range cases {
		inputStr := fmt.Sprintf("%v", c.intervals)
		out := merge(c.intervals)
		fmt.Printf("Input: intervals = %v\nOutput: %v\n\n", inputStr, out)
	}

}

// leetcode start

func merge(intervals [][]int) [][]int {
	out := [][]int{}
	if len(intervals) == 0 {
		return out
	}

	sortIntervals(intervals, 0, len(intervals)-1)
	out = append(out, intervals[0])
	for i, startEnd := range intervals {
		if i == 0 {
			continue
		}

		endIdx := len(out) - 1
		if startEnd[1] <= out[endIdx][1] {
			continue
		}
		if startEnd[0] <= out[endIdx][1] {
			out[endIdx][1] = startEnd[1]
		} else {
			out = append(out, startEnd)
		}
	}

	return out
}

func partition(intervals [][]int, i, j int) int {
	pivot := intervals[i]
	for i < j {
		for i < j && intervals[j][0] >= pivot[0] {
			j--
		}
		intervals[i] = intervals[j]

		for i < j && intervals[i][0] <= pivot[0] {
			i++
		}
		intervals[j] = intervals[i]
	}
	intervals[i] = pivot
	return i
}

func sortIntervals(intervals [][]int, i, j int) {
	if i < j {
		pivotIndex := partition(intervals, i, j)

		sortIntervals(intervals, i, pivotIndex-1)
		sortIntervals(intervals, pivotIndex+1, j)
	}
}
