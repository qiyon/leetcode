class Solution:
	# @param {string} s
	# @return {boolean}
	def isValid(self, s):
		internalStack = []
		for tmpChar in s:
			if internalStack==[] and (tmpChar in ['}',')',']'] ):
				return False
			if internalStack==[] :
				internalStack.append(tmpChar)
				continue
			topChar = internalStack[-1]
			if ( topChar+tmpChar in ['{}','[]','()'] ):
				internalStack.pop()
			else:
				internalStack.append(tmpChar)
		if internalStack == [] :return True
		return False
