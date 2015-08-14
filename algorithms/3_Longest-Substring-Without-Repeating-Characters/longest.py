##leetcode submission begin ---------------------------------
##this solution could do some optimize
class Solution:
    # @param {string} s
    # @return {integer}
    def lengthOfLongestSubstring(self, s):
        charNum = 0;
        d = {}
        for i in range(0, len(s)):
            j = d.get(s[i], -1)
            if (j >= 0):
                d = dict((k, v) for k,v in d.iteritems() if v > j)
            d[s[i]] = i
            nowCount = len(d)
            if nowCount > charNum : charNum = nowCount
        return charNum
##leetcode submission end ---------------------------------

if __name__ == "__main__":
    slObj = Solution();
    longCharNum = slObj.lengthOfLongestSubstring("kqemirokiktoxoikcuuispfjwhqkapymhcaqtpcdzvrzlbloygyqvf")
    print longCharNum;
