package main

import "fmt"

func main() {
	cases := []struct {
		nums []int
		n    int
	}{
		{nums: []int{1, 2, 3, 4, 5}, n: 2},
		{nums: []int{1}, n: 1},
		{nums: []int{1, 2}, n: 1},
		{nums: []int{1, 2}, n: 2},
	}

	for _, c := range cases {
		out := removeNthFromEnd(numsToListNode(c.nums), c.n)
		fmt.Printf("Input: head = %v, n = %d\nOutput: %v\n\n", c.nums, c.n, listNodeToNums(out))
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		p         *ListNode = head
		beforeEnd *ListNode
	)

	for p != nil {
		if n <= 0 {
			if beforeEnd == nil {
				beforeEnd = head
			} else {
				beforeEnd = beforeEnd.Next
			}
		}

		n -= 1
		p = p.Next
	}
	if beforeEnd == nil {
		return head.Next
	} else {
		beforeEnd.Next = beforeEnd.Next.Next
	}
	return head
}
