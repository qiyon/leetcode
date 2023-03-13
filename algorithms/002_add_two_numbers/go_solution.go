package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		l1 []int
		l2 []int
	}{
		{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}},
		{l1: []int{0}, l2: []int{0}},
		{l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}},
	}

	for _, c := range cases {
		l1 := initListNode(c.l1)
		l2 := initListNode(c.l2)

		out := addTwoNumbers(l1, l2)
		fmt.Printf("Input: l1 = %s, l2 = %s \nOutput: %s\n\n", toJsonStr(c.l1), toJsonStr(c.l2), toJsonStr(listNodeToNums(out)))
	}
}

func initListNode(nums []int) *ListNode {
	var p *ListNode
	l := &ListNode{}
	for _, v := range nums {
		if p == nil {
			p = l
		} else {
			p.Next = &ListNode{}
			p = p.Next
		}
		p.Val = v
	}

	return l
}

func toJsonStr(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

func listNodeToNums(l *ListNode) []int {
	var nums []int
	for nil != l {
		nums = append(nums, l.Val)
		l = l.Next
	}

	return nums
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		out       = &ListNode{}
		carry int = 0
		p     *ListNode
	)

	for !(l1 == nil && l2 == nil && carry == 0) {
		if p == nil {
			p = out
		} else {
			p.Next = &ListNode{}
			p = p.Next
		}

		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			p.Val = sum - 10
			carry = 1
		} else {
			p.Val = sum
			carry = 0
		}
	}

	return out
}
