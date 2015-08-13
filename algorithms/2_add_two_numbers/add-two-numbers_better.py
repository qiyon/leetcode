#!/bin/env python
# Definition for singly-linked list.

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


## leetcode submissions begin ----------------------
class Solution:
    # @param {ListNode} l1
    # @param {ListNode} l2
    # @return {ListNode}
    def addTwoNumbers(self, l1, l2):
        cin = 0
        sumHead = None
        sumEnd = None
        while l1 or l2 or cin != 0:
            sumVal = cin
            sumVal += l1.val if l1 else 0
            sumVal += l2.val if l2 else 0
            cin = 1 if sumVal>9 else 0
            if  cin == 1  : sumVal = sumVal - 10
            newNode = ListNode(sumVal)
            if sumHead :
                sumEnd.next = newNode    
                sumEnd = sumEnd.next
            else:
                sumEnd = newNode
                sumHead = sumEnd
            if l1 : l1 = l1.next
            if l2 : l2 = l2.next
        return sumHead
## leetcode submissions end ----------------------

if __name__ == "__main__":
    slObj = Solution()
    a = ListNode(1)
    a.next = ListNode(2)
    b = ListNode(9)
    b.next = ListNode(7)
    c = slObj.addTwoNumbers(a, b)
    print c.val
    print c.next.val
    print c.next.next.val




