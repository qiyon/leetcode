package main

import "fmt"

//leetcode begin
type DpRecord struct {
	src_str string
	dp_map  []int
	src_len int
	max     int
}

func (self *DpRecord) init() {
	self.dp_map[0] = 0
}

func (self *DpRecord) runDp() {
	for index := 1; index < self.src_len; index++ {
		pre_len := self.dp_map[index-1]
		left_index := index - 1 - pre_len
		if self.isParenthesesMatch(left_index, index) {
            self.dp_map[index] = self.dp_map[index-1] + 2
            //for scenes: ()(()) ==> 8
			if left_index-1 >= 0 {
				self.dp_map[index] += self.dp_map[left_index-1]
			}
		} else {
			self.dp_map[index] = 0
		}
		self.updateMax(self.dp_map[index])
	}
}

func (self *DpRecord) updateMax(max_val int) {
	if self.max < max_val {
		self.max = max_val
	}
}

func (self DpRecord) isParenthesesMatch(left_index int, right_index int) bool {
	return left_index >= 0 && self.src_str[left_index] == '(' && self.src_str[right_index] == ')'
}

func longestValidParentheses(s string) int {
	s_len := len(s)
	record := DpRecord{src_str: s, src_len: s_len, max: 0, dp_map: make([]int, s_len)}
	record.init()
	record.runDp()
	return record.max
}

//leetcode end
func main() {
	test_strs := []string{"()", ")(", ")(()())()", "()))"}
	for _, str := range test_strs {
		fmt.Printf("%v \n", str)
		max := longestValidParentheses(str)
		fmt.Printf("%v \n", max)
	}
}
