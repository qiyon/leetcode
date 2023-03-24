package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cases := []struct {
		lists [][]int
	}{
		{lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}},
		{lists: [][]int{}},
		{lists: [][]int{{}}},
	}

	for _, c := range cases {
		var lists []*ListNode
		for _, nums := range c.lists {
			lists = append(lists, numsToListNode(nums))
		}
		out := mergeKLists(lists)
		fmt.Printf("Input: lists = %v\nOutput: %v\n\n", toJsonStr(c.lists), toJsonStr(listNodeToNums(out)))
	}
}

func toJsonStr(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}

func numsToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var (
		head = &ListNode{}
		p    *ListNode
	)
	for i, n := range nums {
		if i == 0 {
			p = head
		} else {
			p.Next = &ListNode{}
			p = p.Next
		}
		p.Val = n
	}

	return head
}

func listNodeToNums(l *ListNode) []int {
	out := []int{}
	if l == nil {
		return out
	}

	p := l
	for nil != p {
		out = append(out, p.Val)
		p = p.Next
	}

	return out
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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	out := lists[0]
	for i := 1; i < len(lists); i++ {
		out = internalMergeTwo(out, lists[i])
	}

	return out
}

func internalMergeTwo(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		head, headEnd, p1, p2 *ListNode
	)

	p1 = list1
	p2 = list2

	for !(p1 == nil && p2 == nil) {
		var appendVal int
		if p1 == nil || (p1 != nil && p2 != nil && p2.Val < p1.Val) {
			appendVal = p2.Val
			p2 = p2.Next
		} else {
			appendVal = p1.Val
			p1 = p1.Next
		}

		if headEnd == nil {
			head = &ListNode{}
			headEnd = head
		} else {
			headEnd.Next = &ListNode{}
			headEnd = headEnd.Next
		}

		headEnd.Val = appendVal
	}

	return head
}
