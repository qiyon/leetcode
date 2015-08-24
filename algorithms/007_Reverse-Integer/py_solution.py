## leetcode submission begin --------------------
class Solution:
	# @param {integer} x
	# @return {integer}
	def reverse(self, x):
		if (x < 0) : 
			isNeg =	True 
		else:
			isNeg = False
		if isNeg :
			x = -x
		ret = 0
		while x > 0:
			before = ret
			ret = (ret * 10) + (x % 10)
			if ret < before:
				return 0
			x = x//10
		if isNeg :
			ret = -ret
		if ret > 2147483647 or ret < -2147483648:
			return 0
		return ret
## leetcode submission end --------------------


if __name__ == "__main__":
    slObj = Solution()
    print slObj.reverse(1231)
