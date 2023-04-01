package main

import (
	"fmt"
)

func main() {
	cases := []int{1, 2, 3, 4, 5, 10}
	for _, c := range cases {
		fmt.Printf("Say %v: %v \n", c, countAndSay(c))
	}
}

// leetcode start

func countAndSay(n int) string {
	result := "1"
	for i := 0; i < n-1; i++ {
		result = say(result)
	}
	return result
}

func say(in string) string {
	out := ""
	arr := []byte(in)
	length := len(arr)
	cnt := 0
	preByte := arr[0]
	for i := 0; i < length; i++ {
		if arr[i] == preByte {
			cnt = cnt + 1
		} else {
			out = out + string(byte('0'+cnt)) + string(preByte)
			preByte = arr[i]
			cnt = 1
		}
	}
	out = out + string(byte('0'+cnt)) + string(preByte)
	return out
}
