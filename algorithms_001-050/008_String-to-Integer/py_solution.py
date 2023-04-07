## leetcode submission begin ---------------------
class Solution:
	# @param {string} str
	# @return {integer}
	def myAtoi(self, str):
		if str == '' :
			return 0
		i = 0
		retInt = 0
		isNeg = False
		yetIntChar = False
		while i < len(str) : 
			theChar = str[i]
			i += 1
			if theChar == ' ' or theChar == '+' or theChar == '-':
				if yetIntChar :
					break
				if theChar == '+' or theChar =='-':
					yetIntChar = True
				if theChar == '-':
					isNeg = ~isNeg 
				continue
			if (ord(theChar) <= ord('9') and ord(theChar) >= ord('0')) :
				yetIntChar = True
				retInt = retInt * 10 + (ord(theChar) - ord('0'))
			else:
				break
			if retInt > 2**31-1 and ~isNeg:
				return 2**31-1
			if retInt > 2**31 and isNeg:
				return -2**31
		if isNeg :
			retInt = -retInt
		return retInt
## leetcode submission end ---------------------

if __name__ == "__main__":
    slObj = Solution()
    print slObj.myAtoi("-1.2")
