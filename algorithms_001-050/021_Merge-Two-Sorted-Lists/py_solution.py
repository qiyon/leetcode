# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
	# @param {ListNode} l1
	# @param {ListNode} l2
	# @return {ListNode}
	def mergeTwoLists(self, l1, l2):
		headNode = None
		lastNode = None
		isAsc = self.isAscList(l1)
		if isAsc == None : isAsc = self.isAscList(l2)
		if isAsc ==None : isAsc = True
		while l1 != None or l2 != None:
			if l1 == None :
				isGetFromL1 = False
			elif l2 == None :
				isGetFromL1 = True
			elif isAsc == None:
				isGetFromL1 = True
			else :
				if isAsc :
					isGetFromL1 = (l1.val <= l2.val)
				else:
					isGetFromL1 = (l1.val >= l2.val)
			if isGetFromL1 :
				pp = l1
				l1 = l1.next
			else:
				pp = l2
				l2 = l2.next
			if lastNode == None:
				headNode = pp
				lastNode = headNode
			else:
				lastNode.next = pp
				lastNode = lastNode.next
		return headNode

	# @param {ListNode} theList
	# @return bool
	def isAscList(self, theList):
		while theList != None:
			if theList.next != None:
				if theList.val < theList.next.val:
					return True
				if theList.val > theList.next.val:
					return False
			theList = theList.next
		return None
