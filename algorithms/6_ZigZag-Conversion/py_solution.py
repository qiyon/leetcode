##leetcode submission begin -----------------
class Solution:		
	# @param {string} s
	# @param {integer} numRows
	# @return {string}
	def convert(self, s, numRows):
		if numRows == 1:
			return s
		if numRows ==0:
			return ""
		outArr = []
		for i in range(0, numRows):
			outArr.append('')
		k = numRows+numRows-2
		for i in range(0,len(s)):
			index = i % k
			if index > numRows-1:
				index = numRows - 1 - (index + 1 - numRows)
			outArr[index] += s[i]
		outStr = ''
		for i in range(0, numRows):
			outStr += outArr[i]
		return outStr
##leetcode submission end -----------------

if __name__ == "__main__":
    slObj = Solution()
    print slObj.convert("PAYPALISHIRING", 3)
