package main

import (
	"fmt"
)

//leetcode begin

func combinationSum2(candidates []int, target int) [][]int {
	t := new(tracking)
	t.SetSortNums(candidates)
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

func (tk *tracking) SetSortNums(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		min_index := i
		for j := i; j < length; j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}
		nums[i], nums[min_index] = nums[min_index], nums[i]
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

//leetcode end

func (tk *tracking) debug() {
	fmt.Printf("%v \n", tk.out)
}

func (t *tmp) debug() {
	fmt.Printf("debug tmp: %v \n", t.list)
}

func main() {
	fmt.Printf("%v \n", combinationSum2([]int{2, 3, 6, 7}, 7))
	fmt.Printf("%v \n", combinationSum2([]int{6, 7}, 5))
	fmt.Printf("%v \n", combinationSum2([]int{}, 10))
	fmt.Printf("%v \n", combinationSum2([]int{1, 2, 3, 4}, 10))
	fmt.Printf("%v \n", combinationSum2([]int{4, 7}, 20))
	fmt.Printf("%v \n", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Printf("%v \n", combinationSum2([]int{1, 1}, 1))
	fmt.Printf("%v \n", combinationSum2([]int{2, 2, 2}, 2))
}
