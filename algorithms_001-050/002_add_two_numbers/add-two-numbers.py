# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
	# @param {ListNode} l1
	# @param {ListNode} l2
	# @return {ListNode}
	def addTwoNumbers(self, l1, l2):
		thein = 0
		isInit = True
		retHead = None
		moveMark = None
		while l1 != None or l2 != None :
			if l1 == None:
				l1 = ListNode(0)
			if l2 == None:
				l2 = ListNode(0)
			sumVal = thein + l1.val + l2.val
			thein = sumVal // 10
			sumVal = sumVal % 10
			if isInit :
				retHead = ListNode(sumVal)
				isInit = False
				moveMark = retHead
			else:
				moveMark.next = ListNode(sumVal)
				moveMark = moveMark.next
			l1 = l1.next
			l2 = l2.next
		else:
			if thein > 0:
				sumVal = thein
				if isInit :
					retHead = ListNode(sumVal)
					isInit = False
					moveMark = retHead
				else:
					moveMark.next = ListNode(sumVal)
					moveMark = moveMark.next
		return retHead