package main

import (
	"fmt"
)

//leetcode begin
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
	pre_byte := arr[0]
	for i := 0; i < length; i++ {
		if arr[i] == pre_byte {
			cnt = cnt + 1
		} else {
			out = out + string(byte('0'+cnt)) + string(pre_byte)
			pre_byte = arr[i]
			cnt = 1
		}
	}
	out = out + string(byte('0'+cnt)) + string(pre_byte)
	return out
}
//leetcode end

func main() {
	fmt.Printf("Say %v: %v \n", 1, countAndSay(1))
	fmt.Printf("Say %v: %v \n", 2, countAndSay(2))
	fmt.Printf("Say %v: %v \n", 3, countAndSay(3))
	fmt.Printf("Say %v: %v \n", 4, countAndSay(4))
	fmt.Printf("Say %v: %v \n", 5, countAndSay(5))
	fmt.Printf("Say %v: %v \n", 10, countAndSay(10))
}
