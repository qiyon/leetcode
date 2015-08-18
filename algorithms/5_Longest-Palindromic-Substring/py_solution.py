
##leetcode submission begin -----------------------
class Solution:
    # @param {string} s
    # @return {string}
    def longestPalindrome(self, s):
        if len(s) <= 1 : return s
        subStrBegin = 0
        subStrEnd = 0
        maxLen = 0
        for i in range(0, len(s) - 1):
            type1 = 1 ## has signle mid
            type2 = 1 ## all double
            j = i + 1
            while type1 or type2 :
                leftIndex = i + i - j
                ##print s
                ##print s[subStrBegin:subStrEnd + 1] , leftIndex ,j 
                ##print "------------------"
                if type1 and leftIndex>=0 :
                    if s[j] == s[leftIndex]:
                        nowLength = j - leftIndex + 1
                        if nowLength > maxLen :
                            maxLen = nowLength
                            subStrBegin = leftIndex
                            subStrEnd = j
                    else:
                        type1 = 0
                if type2 :
                    if s[j] == s[leftIndex + 1]:
                        nowLength = j - leftIndex
                        if nowLength > maxLen:
                            maxLen = nowLength
                            subStrBegin = leftIndex + 1
                            subStrEnd = j
                    else:
                        type2 = 0

                if j==len(s)-1 or leftIndex+1 == 0 :
                    break
                else:
                    j = j + 1
        return s[subStrBegin:subStrEnd + 1] 
    
##leetcode submission end -----------------------


if __name__ == "__main__":
    slObj = Solution();
    print slObj.longestPalindrome("one solution, but will timeout")
