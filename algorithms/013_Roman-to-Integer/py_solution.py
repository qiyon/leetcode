##leetcode submission begin -----------------
class Solution:
	# @param {string} s
	# @return {integer}
	def romanToInt(self, s):
		retInt = 0
		if s == '' : return 0
		d = {
			'I' : 1,
			'V' : 5,
			'X' : 10,
			'L' : 50,
			'C' : 100,
			'D' : 500,
			'M' : 1000
		}
		if (len(s) == 1) : return d.get(s , 0)
		for i in range(1, len(s)):
			pre = d.get(s[i-1])
			later = d.get(s[i])
			if pre == None or later == None: return 0
			if i == 1: retInt = pre
			if pre >= later:
				retInt += later
			else:
				retInt += later - 2 * pre
		return retInt
##leetcode submission end -----------------


if __name__ == "__main__":
    sl = Solution()
    print sl.romanToInt("XXX")
