package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		num int
	}{
		{num: 3},
		{num: 58},
		{num: 1994},
		{num: 900},
	}

	for _, c := range cases {
		out := intToRoman(c.num)
		fmt.Printf("Input: num = %d\nOutput: \"%s\"\n\n", c.num, out)
	}
}

// leetcode start

func intToRoman(num int) string {
	var out = ""

	symbolMap := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	loop := []int{1000, 100, 10, 1}

	for i, lv := range loop {
		if num == 0 {
			break
		}

		v := num / lv
		num = num % lv

		if v >= 9 && i > 0 {
			out += symbolMap[lv] + symbolMap[lv*10]
		} else if v >= 4 && i > 0 {
			if v == 4 {
				out += symbolMap[lv]
			}
			out += symbolMap[lv*5]
			r := v - 5
			for r > 0 {
				out += symbolMap[lv]
				r--
			}
		} else if v > 0 {
			r := v
			for r > 0 {
				out += symbolMap[lv]
				r--
			}
		}
	}

	return out
}
