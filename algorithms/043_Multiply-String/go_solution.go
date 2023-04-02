package main

import (
	"fmt"
)

func main() {
	cases := [][]string{
		{"2", "3"},
		{"123", "456"},
		{"12232300", "450057"},
		{"122", "450051237"},
		{"0", "0"},
		{"1", "1"},
		{"9133", "0"},
		{"0", "9133"},
		{"10001", "10000001"},
	}
	for _, c := range cases {
		fmt.Printf("Input: num1 = \"%s\", num2 = \"%s\"\nOutput: \"%s\"\n\n", c[0], c[1], multiply(c[0], c[1]))
	}
}

// leetcode start

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	c := &calcStat{
		num1:    []byte(num1),
		length1: len(num1),
		num2:    []byte(num2),
		length2: len(num2),
		res:     &result{},
	}

	var ci int
	var baseIndex int
	for i := c.length1 - 1; i >= 0; i-- {
		ci = 0
		for j := c.length2 - 1; j >= 0; j-- {
			baseIndex = c.length1 - 1 - i + c.length2 - 1 - j
			ci = c.res.addVal(baseIndex, int(c.num1[i]-'0')*int(c.num2[j]-'0')+ci)
		}
		if ci > 0 {
			c.res.addVal(baseIndex+1, ci)
		}
	}
	return c.res.getValString()
}

type calcStat struct {
	num1    []byte
	num2    []byte
	length1 int
	length2 int
	res     *result
}

type result struct {
	list      []int
	theLength int
}

func (r *result) addVal(baseIndex int, val int) int {
	if r.theLength-1 < baseIndex {
		r.list = append(r.list, val%10)
		r.theLength++
		return val / 10
	} else {
		outCi := (r.list[baseIndex] + val) / 10
		r.list[baseIndex] = (r.list[baseIndex] + val) % 10
		return outCi
	}
}

func (r *result) getValString() string {
	str := ""
	for _, v := range r.list {
		str = string(byte('0'+v)) + str
	}
	return str
}
