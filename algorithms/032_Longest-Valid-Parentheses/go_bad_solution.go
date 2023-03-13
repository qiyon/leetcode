package main

import "fmt"

//leetcode begin
type info struct {
	length int
	need   int
}

func longestValidParentheses_bad(s string) int {
	s_len := len(s)
	if s_len <= 1 {
		return 0
	}
	list := make([]info, s_len)
	for i := 0; i < s_len-1; i++ {
		var tmp_need int
		if s[i] == '(' {
			tmp_need = 1
		} else {
			tmp_need = -1
		}
		list[i] = info{0, tmp_need}
	}
	max := 0
	for end := 1; end < s_len; end++ {
		tmp_char := s[end]
		for i := 0; i < end; i++ {
			if list[i].need == -1 {
				continue
			}
			if tmp_char == '(' {
				list[i].need++
			}
			if tmp_char == ')' {
				list[i].need--
				if list[i].need >= 0 {
					list[i].length += 2
				}
			}
			if list[i].need == 0 && list[i].length > max {
				max = list[i].length
			}
		}
	}
	return max
}

//leetcode end
func main() {
	str := ")()(()())("
	max := longestValidParentheses_bad(str)
	fmt.Printf("%v", max)
}
