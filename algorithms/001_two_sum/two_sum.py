class Solution:
    # @param {integer[]} nums
    # @param {integer} target
    # @return {integer[]}
    def twoSum(self, nums, target):
    d = {}
    for i in range(0 , len(nums)):
        item = nums[i]
        j = d.get(target-item)
        if j > 0:
            return [j , i+1]
        d[item] = i+1
