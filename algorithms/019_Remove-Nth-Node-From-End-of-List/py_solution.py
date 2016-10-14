# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
	# @param {ListNode} head
	# @param {integer} n
	# @return {ListNode}
	def removeNthFromEnd(self, head, n):
		if n == 0 :return head
		befDelNode = head
		lastNode = head
		while lastNode != None :
			lastNode = lastNode.next
			if n <0:
				befDelNode = befDelNode.next
			n -= 1
		# if befDelNode == head : return head.next
		if n >= 0 : return head.next
		befDelNode.next = befDelNode.next.next
		return head
