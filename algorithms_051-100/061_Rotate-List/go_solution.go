package main

import "fmt"

func main() {
	cases := []struct {
		head []int
		k    int
	}{
		{head: []int{1, 2, 3, 4, 5}, k: 2},
		{head: []int{0, 1, 2}, k: 4},
	}

	for _, c := range cases {
		out := rotateRight(numsToListNode(c.head), c.k)
		fmt.Printf("Input: head = %v, k = %d\nOutput: %v\n\n", c.head, c.k, listNodeToNums(out))
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	listLen := 1
	p := head
	var pre *ListNode
	for p.Next != nil {
		pre = p
		p = p.Next
		listLen++
	}

	k = k % listLen
	if k == 0 {
		return head
	}

	mv := listLen - k + 1
	p.Next = head
	for mv > 0 {
		p = p.Next
		pre = pre.Next
		mv--
	}

	pre.Next = nil
	return p
}
