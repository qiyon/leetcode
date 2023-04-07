class Solution:
	# @param {integer[]} nums
	# @return {integer}
	def removeDuplicates(self, nums):
		if nums == [] : return 0
		d = {}
		count = 0
		i = 0
		while i < len(nums):
			if d.get(nums[i]) == None:
				count += 1
				d[nums[i]] = 1
				i += 1
			else:
				del nums[i]
		print nums
		return count
