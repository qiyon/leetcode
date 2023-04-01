package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		candidates []int
		target     int
	}{
		{candidates: []int{2, 3, 6, 7}, target: 7},
		{candidates: []int{6, 7}, target: 5},
		{candidates: []int{}, target: 10},
		{candidates: []int{1, 2, 3, 4}, target: 10},
		{candidates: []int{4, 7}, target: 20},
		{candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 10},
		{candidates: []int{1, 1}, target: 1},
		{candidates: []int{2, 2, 2}, target: 2},
	}
	for _, c := range cases {
		out := combinationSum2(c.candidates, c.target)
		fmt.Printf("Input: candidates = %v, target = %d\nOutput: %s\n\n", c.candidates, c.target, toJsonStr(out))
	}
}

func (tk *tracking) debug() {
	fmt.Printf("%v \n", tk.out)
}

func (t *tmp) debug() {
	fmt.Printf("debug tmp: %v \n", t.list)
}

func toJsonStr(in interface{}) string {
	out, _ := json.Marshal(in)
	return string(out)
}

// leetcode start

func combinationSum2(candidates []int, target int) [][]int {
	t := &tracking{
		target:     target,
		numsLength: len(candidates),
		out:        [][]int{},
	}
	t.SetSortNums(candidates)

	tTemp := new(tmp)
	tTemp.cnt = 0
	t.findResult(tTemp, 0)
	return t.out
}

type tracking struct {
	nums       []int
	numsLength int
	target     int
	out        [][]int
}

func (tk *tracking) SetSortNums(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	tk.nums = nums
}

func (tk *tracking) findResult(t *tmp, numsStart int) {
	for index := numsStart; index < tk.numsLength; index++ {
		v := tk.nums[index]
		loopTmp := newTmpByCopy(t)
		if loopTmp.cnt+v < tk.target {
			newTemp := newTmpByOldAppend(loopTmp, v)
			tk.findResult(newTemp, index+1)
		}
		if loopTmp.cnt+v == tk.target {
			tk.out = append(tk.out, append(loopTmp.list, v))
		}
		for index+1 < tk.numsLength && tk.nums[index] == tk.nums[index+1] {
			index++
		}
	}
}

type tmp struct {
	list []int
	cnt  int
}

func newTmpByOldAppend(srcTemp *tmp, n int) *tmp {
	newTemp := new(tmp)
	newTemp.list = append(srcTemp.list, n)
	newTemp.cnt = srcTemp.cnt + n
	return newTemp
}

func newTmpByCopy(srcTemp *tmp) *tmp {
	newTemp := new(tmp)
	newTemp.list = make([]int, len(srcTemp.list))
	copy(newTemp.list, srcTemp.list)
	newTemp.cnt = srcTemp.cnt
	return newTemp
}
