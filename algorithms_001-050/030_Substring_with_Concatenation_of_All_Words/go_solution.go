package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		s     string
		words []string
	}{
		{s: "barfoothefoobarman", words: []string{"foo", "bar"}},
		{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}},
		{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}},
	}
	for _, c := range cases {
		out := findSubstring(c.s, c.words)
		fmt.Printf("s: %s, words:%v \nOutput:%v\n\n", c.s, c.words, out)
	}
}

// leetcode start

func findSubstring(s string, words []string) []int {
	indexes := []int{}

	wordNum := len(words)
	if wordNum == 0 {
		return indexes
	}
	strLen := len(s)
	wordLen := len(words[0])
	if strLen == 0 || wordLen == 0 {
		return indexes
	}

	wordMap := make(map[string]int)
	for _, w := range words {
		cnt, ok := wordMap[w]
		if ok {
			wordMap[w] = cnt + 1
		} else {
			wordMap[w] = 1
		}
	}

	start := 0
	end := 0
	for start < strLen && start+wordLen*wordNum <= strLen && end < strLen {
		markMap := make(map[string]int)
		i := 0
		for i < wordNum {
			iWord := s[start+wordLen*i : start+wordLen*(i+1)]
			wordCnt, ok := wordMap[iWord]
			markCnt, markOk := markMap[iWord]
			if !ok || (ok && markOk && wordCnt <= markCnt) {
				end = start + wordLen*wordNum
				break
			} else {
				if markOk {
					markMap[iWord] = markCnt + 1
				} else {
					markMap[iWord] = 1
				}
			}
			i++
		}

		if i == wordNum {
			if matchMap(wordMap, markMap) {
				indexes = append(indexes, start)
				start++
				continue
			} else {
				end = start + wordLen*wordNum
			}
		}
		for end < strLen {
			endWord := s[end-wordLen+1 : end+1]
			_, ok := wordMap[endWord]
			if ok {
				start++
				break
			} else {
				end++
			}
		}
	}
	return indexes
}

func matchMap(wordMap map[string]int, markMap map[string]int) bool {
	for w, cnt := range wordMap {
		markCnt, markOk := markMap[w]
		if !markOk {
			return false
		}
		if markCnt != cnt {
			return false
		}
	}
	return true
}
