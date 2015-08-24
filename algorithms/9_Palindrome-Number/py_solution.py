##leetcode submission begin----------------
class Solution:
	# @param {integer} x
	# @return {boolean}
	def isPalindrome(self, x):
		if x<0 : return False
		if x>=0 and x<= 9 : return True
		a = 0
		while x > 0 :
			b = x % 10
			if a==0 and b == 0 :return False
			a = a * 10 + b
			if x == a : return True
			x = x // 10
			if x == a : return True
			if x<a : return False
		return False
##leetcode submission end----------------


if __name__ == "__main__":
    slObj = Solution()
    print slObj.isPalindrome(525)
