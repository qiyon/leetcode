class Solution:
	# @param {integer[]} nums
	# @param {integer} val
	# @return {integer}
	def removeElement(self, nums, val):
		if nums == [] :return 0
		i = 0
		for j in range(0, len(nums)):
			if nums[i] == val:
				del nums[i]
			else:
				i += 1
		return len(nums)
