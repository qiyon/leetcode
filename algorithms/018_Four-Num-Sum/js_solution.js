/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[][]}
 */
var fourSum = function(nums, target) {
    var out = [];
    var len = nums.length;
    var i, j, left, right, tmp_target;
    nums = nums.sort(function (a, b) {
      return a - b;
    });
    for (i = 0; i < len - 3; i++) {
      if (i > 0 && nums[i] == nums[i - 1]){
        continue;
      }
      for (j = i + 1; j < len -2; j++) {
        if (j > i + 1 && nums[j] == nums[j - 1]){
          continue;
        }
        left = j + 1;
        right = len - 1;
        while (left < right) {
          tmp_target = nums[i] + nums[j] + nums[left] + nums[right];
          if (tmp_target < target) {
            left++;
          } else if (tmp_target > target){
            right--;
          }else{
            out.push([nums[i], nums[j], nums[left], nums[right]]);
            while (left < right && nums[left] == nums[left + 1]){
              left++;
            }
            left++;
          }
        }
      }
    }
    return out;
};

console.log(fourSum([-3, -2, -1, 0, 0, 1, 2, 3], 0));
