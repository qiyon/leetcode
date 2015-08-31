##leetcode submission begin -----------------------
class Solution:
	# @param {string[]} strs
	# @return {string}
	def longestCommonPrefix(self, strs):
		longPre = ''
		if strs == [] : return ''
		i = 0
		while True:
			tmpChar = None
			for inStr in strs:
				if i < len(inStr) :
					getChar = inStr[i] 
				else:
					tmpChar = None
					break
				if tmpChar == None : tmpChar = getChar
				if getChar != tmpChar:
					tmpChar = None
					break
			if tmpChar == None :break
			longPre = longPre + tmpChar
			i += 1
		return longPre
##leetcode submission end -----------------------

if __name__ == "__main__":
    slObj = Solution()
    print slObj.longestCommonPrefix(["abcd","ab"])
