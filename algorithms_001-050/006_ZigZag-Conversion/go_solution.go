package main

import "fmt"

func main() {
	cases := []struct {
		s       string
		numRows int
	}{
		{s: "PAYPALISHIRING", numRows: 3},
		{s: "PAYPALISHIRING", numRows: 4},
		{s: "A", numRows: 1},
	}

	for _, c := range cases {
		out := convert(c.s, c.numRows)
		fmt.Printf("Input: s = %s, numRows = %d\nOutput: %s \n\n", c.s, c.numRows, out)
	}

}

// leetcode start

func convert(s string, numRows int) string {
	var (
		lines           = make([][]rune, numRows)
		arrIndex    int = 0
		indexOffset int = 1
	)

	if s == "" || numRows <= 1 {
		return s
	}

	for _, char := range s {
		lines[arrIndex] = append(lines[arrIndex], char)
		arrIndex += indexOffset
		if arrIndex == numRows-1 {
			indexOffset = -1
		}
		if arrIndex == 0 {
			indexOffset = 1
		}
	}

	out := ""
	for _, line := range lines {
		out += string(line)
	}
	return out
}
