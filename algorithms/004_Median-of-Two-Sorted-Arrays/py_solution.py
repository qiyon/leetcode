## leetcode submission begin ---------------------
class Solution:
    # @param {integer[]} nums1
    # @param {integer[]} nums2
    # @return {float}
    def findMedianSortedArrays(self, nums1, nums2):
        delNum = (len(nums1) + len(nums2) + 1 ) / 2 - 1                       
        if delNum == -1 : return 0
        justOne = (len(nums1) + len(nums2)) % 2
        getNumList = []
        while delNum >= 0  :
            ##print delNum , justOne , len(nums1) ,len(nums2)
            #print nums1
            #print nums2
            #print getNumList
            if (delNum == 0):
                if (len(nums1) != 0 and len(nums2) != 0):
                    getNumList.append(min(nums1[0], nums2[0]))
                else:
                    if len(nums1) == 0 : getNumList.append(nums2[0])
                    if len(nums2) == 0 : getNumList.append(nums1[0])
                if justOne :
                    delNum = -1
                else:
                    delNum = 1
                    justOne = 1
                continue
            if (len(nums1) != 0 and len(nums2) != 0 ):
                toDel1 = (delNum + 1) / 2
                toDel2 = toDel1
                if toDel1 > len(nums1) : toDel1 = len(nums1)
                if toDel2 > len(nums2) : toDel2 = len(nums2)
                if nums1[toDel1 - 1] > nums2[toDel2 - 1] :
                    nums2 = nums2[toDel2:]
                    delNum = delNum - toDel2
                else:
                    nums1 = nums1[toDel1:]
                    delNum = delNum - toDel1
            else:
                if len(nums1) == 0 :
                    nums2 = nums2[delNum:]
                if len(nums2) == 0 :
                    nums1 = nums1[delNum:]
                delNum = 0
        #print getNumList
        #print "------------"
        if len(getNumList) == 2:
            return (getNumList[0] + getNumList[1]) / 2.0
        else:
            return getNumList[0]

##leetcode submission end ---------------------

if __name__ == "__main__":
    slObj = Solution()
    ##print slObj.findMedianSortedArrays([1,2,5,7,9,12,16,18,23], [1,2,3,4,5,6,6.5,8,9])
    ##print slObj.findMedianSortedArrays([1,2,3,4,5,6,7,8],[])
    print slObj.findMedianSortedArrays([1,2,3,4,5,6,7,8],[100])
