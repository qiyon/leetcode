package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		nums []int
	}{
		{nums: []int{1, 1, 2}},
		{nums: []int{1, 2, 3}},
		{nums: []int{}},
		{nums: []int{1}},
		{nums: []int{1, 1, 1, 1, 1, 1}},
		{nums: []int{1, 1, 1, 1, 2, 1, 1, 1, 1}},
		{nums: []int{1, 2, 3, 4}},
		{nums: []int{3, 6, 9}},
	}
	for _, c := range cases {
		out := permuteUnique(c.nums)
		fmt.Printf("Input: nums = %v\nOutput: %v\n\n", c.nums, out)
	}
}

// leetcode start

func permuteUnique(nums []int) [][]int {
	c := new(calc)
	c.theLength = len(nums)
	//may use sort.Ints()
	qsort(nums)
	//fmt.Printf("nums: %v \n", nums)
	c.nums = nums
	if c.theLength == 0 {
		return [][]int{}
	}
	c.out = make([][]int, 0)
	tmp := new(outTmp)
	tmp.list = make([]int, c.theLength)
	tmp.theLength = 0
	c.permuteAll(tmp, nums)
	return c.out
}

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}

type calc struct {
	nums      []int
	theLength int
	out       [][]int
}

func (c *calc) permuteAll(t *outTmp, left []int) {
	if t.theLength == c.theLength {
		c.addOut(t)
	}
	for i := range left {
		if i > 0 && left[i] == left[i-1] {
			continue
		}
		newLeft := c.copyNewLeftByDeleteOne(left, c.theLength-t.theLength-1, i)
		//fmt.Printf("remove %v: %v\n", i, newLeft)
		t.add(left[i])
		c.permuteAll(t, newLeft)
		t.pop()
	}
}

func (c *calc) copyNewLeftByDeleteOne(preLeft []int, newLeftLen int, deleteIndex int) []int {
	newLeft := make([]int, newLeftLen)
	markIndex := 0
	for i := range preLeft {
		if i == deleteIndex {
			continue
		}
		newLeft[markIndex] = preLeft[i]
		markIndex++
	}
	return newLeft
}

func (c *calc) addOut(t *outTmp) {
	//fmt.Printf("%v\n", t.list)
	res := make([]int, c.theLength)
	for i := 0; i < c.theLength; i++ {
		res[i] = t.list[i]
	}
	c.out = append(c.out, res)
}

type outTmp struct {
	list      []int
	theLength int
}

func (t *outTmp) add(val int) {
	t.list[t.theLength] = val
	t.theLength++
}

func (t *outTmp) pop() {
	t.theLength--
}
