package main

import "fmt"

func main() {
	cases := []struct {
		list1 []int
		list2 []int
	}{
		{list1: []int{1, 2, 4}, list2: []int{1, 3, 4}},
		{list1: []int{}, list2: []int{}},
		{list1: []int{}, list2: []int{0}},
	}

	for _, c := range cases {
		out := mergeTwoLists(numsToListNode(c.list1), numsToListNode(c.list2))
		fmt.Printf("Input: list1 = %v, list2 = %v\nOutput: %v\n\n", c.list1, c.list2, listNodeToNums(out))
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
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

// leetcode start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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
