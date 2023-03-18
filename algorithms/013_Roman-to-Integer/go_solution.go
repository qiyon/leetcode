package main

import "fmt"

func main() {
	cases := []struct {
		s      string
		should int
	}{
		{s: "III", should: 3},
		{s: "LVIII", should: 58},
		{s: "MCMXCIV", should: 1994},
	}

	for _, c := range cases {
		out := romanToInt(c.s)
		fmt.Printf("Input: s = \"%s\" \nShould: %d \nOutput: %d\n\n", c.s, c.should, out)
	}
}

// leetcode start

func romanToInt(s string) int {
	var romanMap = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var (
		out             = 0
		beforeChar rune = '0'
		beforeSum  int
	)
	for _, v := range s {
		num := romanMap[v]
		out += num

		if beforeChar == v {
			beforeSum += num
		} else if beforeChar != '0' && romanMap[v] > romanMap[beforeChar] {
			// 罗马数字前者比后者大，减去已经相加的值
			out -= 2 * beforeSum
			beforeSum = 0
		} else {
			beforeChar = v
			beforeSum = num
		}
	}

	return out
}
