/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 **/
var threeSumClosest = function(nums, target) {
    var length = nums.length;
    nums = nums.sort(function (a, b){
        return a - b;
    });
    // console.log(nums);
    if (length < 3){
        return null;
    }
    result = nums[0] + nums[1] + nums[2];
    if (result == target ) {
        return result;
    }
    var now_closest = Math.abs(result - target);
    var left, right;
    var in_closest;
    for (var i = 0; i < length - 2; i++) {
        //remove first repeat
        if (i > 0 && nums[i] == nums[i - 1]){
            continue;
        }
        left = i + 1;
        right = length - 1;
        while (left < right) {
            in_closest = nums[i] + nums[left] + nums[right] - target;
            // console.log('---------------');
            // console.log(i, left, right);
            // console.log(nums[i], nums[left], nums[right]);
            // console.log(nums[i] + nums[left] + nums[right], in_closest, now_closest);
            if (in_closest <= (-1 * now_closest)){
                left++;
            }else if (in_closest >= now_closest){
                right--;
            }else{
                // console.log('new less');
                result = nums[i] + nums[left] + nums[right];
                if (result == target){
                    return result;
                }
                now_closest = Math.abs(result - target);
                //remove second repeat
                while (left < right && nums[left] == nums[left + 1]){
                    left++;
                }
                if (in_closest < 0){
                  left++;
                }else{
                  right--;
                }
            }
        }
    }
    return result;
};


console.log(threeSumClosest([-1, 2, 1, -4], 1));
console.log(threeSumClosest([-1, 2, 1, -4, 5, 6, 10, 101, 1000], 100));
console.log(threeSumClosest([100,1,-3,3,5,4,1], 1));
