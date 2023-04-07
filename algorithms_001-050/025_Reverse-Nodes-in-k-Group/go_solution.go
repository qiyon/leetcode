package main

import "fmt"

func main() {
	cases := []struct {
		head []int
		k    int
	}{
		{head: []int{1, 2, 3, 4, 5}, k: 3},
		{head: []int{1, 2, 3, 4, 5, 6}, k: 3},
		{head: []int{1, 2, 3}, k: 3},
		{head: []int{1, 2, 3, 4}, k: 3},
		{head: []int{}, k: 2},
		{head: []int{1, 2}, k: 3},
	}

	for _, c := range cases {
		out := reverseKGroup(numsToListNode(c.head), c.k)
		fmt.Printf("Input: head = %v, k = %v\nOutput: %v\n\n", c.head, c.k, listNodeToNums(out))
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	if head == nil || head.Next == nil {
		return head
	}

	fixedHead := &ListNode{Next: head}
	pre := fixedHead
	for pre.Next != nil {
		p := pre.Next

		// check
		cnt := 0
		cntPointer := p
		for cntPointer != nil {
			cnt++
			cntPointer = cntPointer.Next
		}
		if cnt < k {
			break
		}

		// swap move, for k = 3:
		//      p   q
		//      \   /
		// 3-2-1-4-5-6
		//     |
		//    pre
		for i := 2; i <= k; i++ {
			q := p.Next
			p.Next = q.Next
			q.Next = pre.Next
			pre.Next = q
		}

		pre = p
	}

	return fixedHead.Next
}
