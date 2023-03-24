package main

import "fmt"

func main() {
	cases := []struct {
		head []int
	}{
		{head: []int{1, 2, 3, 4}},
		{head: []int{}},
		{head: []int{1}},
		{head: []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		out := swapPairs(numsToListNode(c.head))
		fmt.Printf("Input: head = %v\nOutput: %v\n\n", c.head, listNodeToNums(out))
	}
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	for p.Next != nil {
		if p == head {
			head = p.Next
			p.Next = head.Next
			head.Next = p
			continue
		}

		q := p.Next
		if q.Next == nil {
			break
		}

		// 2-1-3-4-5
		//   | |
		//   p q
		p.Next = q.Next
		q.Next = q.Next.Next
		p.Next.Next = q
		p = q
	}

	return head
}
