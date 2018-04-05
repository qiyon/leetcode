package main

import "fmt"

//leetcode begin

func permute(nums []int) [][]int {
	c := new(calc)
	c.nums = nums
	c.theLength = len(nums)
	if c.theLength == 0 {
		return [][]int{}
	}
	resCount := 1
	for i := 2; i <= c.theLength; i++ {
		resCount = resCount * i
	}
	c.out = make([][]int, resCount)
	c.yetOut = 0
	tmp := new(outTmp)
	tmp.list = make([]int, c.theLength)
	tmp.theLength = 0
	c.permuteAll(tmp, nums)
	return c.out
}

type calc struct {
	nums      []int
	theLength int
	out       [][]int
	yetOut    int
}

func (c *calc) permuteAll(t *outTmp, left []int) {
	if t.theLength == c.theLength {
		c.addOut(t)
	}
	for i := range left {
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
	c.out[c.yetOut] = make([]int, c.theLength)
	for i := 0; i < c.theLength; i++ {
		c.out[c.yetOut][i] = t.list[i]
	}
	c.yetOut++
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

//leetcode end

func main() {
	fmt.Printf("%v \n", permute([]int{}))
	fmt.Printf("%v \n", permute([]int{1}))
	fmt.Printf("%v \n", permute([]int{1, 2, 3}))
	fmt.Printf("%v \n", permute([]int{1, 2, 3, 4}))
	fmt.Printf("%v \n", permute([]int{3, 6, 9}))
}
