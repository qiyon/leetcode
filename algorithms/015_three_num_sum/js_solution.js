//-----leetcode sbumit-----

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    nums = nums.sort(function(a, b) {
        return a - b;
    });
    var i, j, k;
    var length = nums.length;
    var left, right;
    var sum, target;
    var result = [];
    for(i = 0; i < length - 2; i++){
        //remove first repeat
        if (i > 0 && nums[i] == nums[i - 1]){
            continue;
        }
        left = i + 1;
        right = length - 1;
        target = -1 * nums[i];
        while (left < right){
            sum = nums[left] + nums[right];
            if (sum < target){
                left++;
            }else if (sum > target){
                right--;
            }else{
                result.push([nums[i], nums[left], nums[right]]);
                //remove second repeat
                while (left < right && nums[left] == nums[left + 1]){
                    left++;
                }
                left++;
            }
        }
    }
    return result;
};
//-----leetcode sbumit-----


var S = [-1, 0, 1, 2, -1, -4];
var result = threeSum(S);
console.log(result);
console.log(threeSum([-1, -1, -1, -1, -1, -1, 2, 3, 4]))
console.log(threeSum([-1, 1, -1, 1, -1, -1, 2, -3, 3, 4]))
console.log(threeSum([-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]))
