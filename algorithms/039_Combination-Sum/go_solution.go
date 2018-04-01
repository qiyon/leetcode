package main

import (
	"fmt"
)

//leetcode begin

func combinationSum(candidates []int, target int) [][]int {
	t := new(tracking)
	t.nums = candidates
	t.target = target
	t.numsLength = len(candidates)
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

func (tk *tracking) findResult(t *tmp, numsStart int) {
	for index := numsStart; index < tk.numsLength; index++ {
		v := tk.nums[index]
		loopTmp := newTmpByCopy(t)
		if loopTmp.cnt+v == tk.target {
			tk.out = append(tk.out, append(loopTmp.list, v))
			continue
		}
		if loopTmp.cnt+v > tk.target {
			continue
		}
		newTemp := newTmpByOldAppend(loopTmp, v)
		tk.findResult(newTemp, index)
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

//leetcode end

func (tk *tracking) debug() {
	fmt.Printf("%v \n", tk.out)
}

func (t *tmp) debug() {
	fmt.Printf("debug tmp: %v \n", t.list)
}

func main() {
	fmt.Printf("%v \n", combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Printf("%v \n", combinationSum([]int{6, 7}, 5))
	fmt.Printf("%v \n", combinationSum([]int{}, 10))
	fmt.Printf("%v \n", combinationSum([]int{2, 3, 7}, 20))
	fmt.Printf("%v \n", combinationSum([]int{2, 3, 4}, 10))
	fmt.Printf("%v \n", combinationSum([]int{2, 7}, 20))
}
