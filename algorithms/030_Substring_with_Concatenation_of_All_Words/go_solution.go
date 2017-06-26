package main

import "fmt"

//--leetcode begin
func findSubstring(s string, words []string) []int {
	word_num := len(words)
	if word_num == 0 {
		return []int{}
	}
	s_len := len(s)
	word_len := len(words[0])
	if s_len == 0 || word_len == 0 {
		return []int{}
	}
	idxes := []int{}
	wd_map := make(map[string]int)
	for _, wd := range words {
		yet, ok := wd_map[wd]
		if ok {
			wd_map[wd] = yet + 1
		} else {
			wd_map[wd] = 1
		}
	}
	start := 0
	end := 0
	for start < s_len && start+word_len*word_num <= s_len && end < s_len {
		mark_map := make(map[string]int)
		i := 0
		for i < word_num {
			tmp_wd := s[start+word_len*i:start+word_len*(i+1)]
			cnt, ok := wd_map[tmp_wd]
			mark_cnt, mark_ok := mark_map[tmp_wd]
			if !ok || (ok && mark_ok && cnt <= mark_cnt) {
				end = start + word_len*word_num;
				break
			} else {
				if mark_ok {
					mark_map[tmp_wd] = mark_cnt + 1
				} else {
					mark_map[tmp_wd] = 1
				}
			}
			i++
		}
		if i == word_num {
			if (matchMap(wd_map, mark_map)) {
				idxes = append(idxes, start)
				start++
				continue
			} else {
				end = start + word_len*word_num;
			}
		}
		for end < s_len {
			tmp_wd := s[end-word_len+1:end+1]
			_, ok := wd_map[tmp_wd]
			if ok {
				start++
				break
			} else {
				end++
			}
		}
	}
	return idxes
}

func matchMap(wd_map map[string]int, mark_map map[string]int) bool {
	for wd, cnt := range wd_map {
		mark_cnt, mark_ok := mark_map[wd]
		if !mark_ok {
			return false;
		}
		if mark_cnt != cnt {
			return false;
		}
	}
	return true;
}

//--leetcode end
func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	result := findSubstring(s, words)
	fmt.Printf("%v", result)
}
