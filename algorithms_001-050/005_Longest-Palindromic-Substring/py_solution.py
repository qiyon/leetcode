
##leetcode submission begin -----------------------
class Solution:
    # @param {string} s
    # @return {string}
    def longestPalindrome(self, s):
        subStr = ""
        for i in range(0, len(s)):
            for k in range(0,2):
                tmp = self.getSubStr(s, i, i+k)
                if len(tmp) > len(subStr):
                    subStr = tmp
        return subStr

    def getSubStr(self, s, left, right):
        while left>=0 and right<len(s) and s[left]==s[right]:
            left -= 1
            right += 1
        return s[left+1 : right]

##leetcode submission end -----------------------


if __name__ == "__main__":
    slObj = Solution();
    print slObj.longestPalindrome("good solution")
