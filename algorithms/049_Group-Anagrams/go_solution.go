package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		strs []string
	}{
		{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
		{strs: []string{"", "abc", "abcd", "acb", "bat", "bat"}},
	}
	for _, c := range cases {
		out := groupAnagrams(c.strs)
		fmt.Printf("Input: strs = %v\nOutput: %v\n\n", c.strs, out)
	}
}

// leetcode start

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)
	keyOrder := make([]string, 0)
	for _, row := range strs {
		groupStr := getGroupStr(row)
		_, ok := hashMap[groupStr]
		if !ok {
			hashMap[groupStr] = make([]string, 0)
			keyOrder = append(keyOrder, groupStr)
		}
		hashMap[groupStr] = append(hashMap[groupStr], row)
	}

	res := make([][]string, 0)
	for _, oneGroupStr := range keyOrder {
		srcStrList := hashMap[oneGroupStr]
		res = append(res, srcStrList)
	}
	return res
}

func getGroupStr(str string) string {
	byteList := []byte(str)
	byteQSort(byteList)
	return string(byteList)
}

func byteQSort(data []byte) {
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
	byteQSort(data[:head])
	byteQSort(data[head+1:])
}
