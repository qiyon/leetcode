package main

import "fmt"

func main() {
	cases := []struct {
		s     string
		words []string
	}{
		{s: "barfoothefoobarman", words: []string{"foo", "bar"}},
	}
	for _, c := range cases {
		out := findSubstring(c.s, c.words)
		fmt.Printf("s: %s, words:%v \n %v\n\n", c.s, c.words, out)
	}
}

// leetcode start

func findSubstring(s string, words []string) []int {
	wordNum := len(words)
	if wordNum == 0 {
		return []int{}
	}
	sLen := len(s)
	wordLen := len(words[0])
	if sLen == 0 || wordLen == 0 {
		return []int{}
	}
	idxes := []int{}
	wdMap := make(map[string]int)
	for _, wd := range words {
		yet, ok := wdMap[wd]
		if ok {
			wdMap[wd] = yet + 1
		} else {
			wdMap[wd] = 1
		}
	}
	start := 0
	end := 0
	for start < sLen && start+wordLen*wordNum <= sLen && end < sLen {
		markMap := make(map[string]int)
		i := 0
		for i < wordNum {
			tmpWd := s[start+wordLen*i : start+wordLen*(i+1)]
			cnt, ok := wdMap[tmpWd]
			markCnt, markOk := markMap[tmpWd]
			if !ok || (ok && markOk && cnt <= markCnt) {
				end = start + wordLen*wordNum
				break
			} else {
				if markOk {
					markMap[tmpWd] = markCnt + 1
				} else {
					markMap[tmpWd] = 1
				}
			}
			i++
		}
		if i == wordNum {
			if matchMap(wdMap, markMap) {
				idxes = append(idxes, start)
				start++
				continue
			} else {
				end = start + wordLen*wordNum
			}
		}
		for end < sLen {
			tmpWd := s[end-wordLen+1 : end+1]
			_, ok := wdMap[tmpWd]
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

func matchMap(wdMap map[string]int, markMap map[string]int) bool {
	for wd, cnt := range wdMap {
		markCnt, markOk := markMap[wd]
		if !markOk {
			return false
		}
		if markCnt != cnt {
			return false
		}
	}
	return true
}
