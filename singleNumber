// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

 

// Example 1:

// Input: nums = [2,2,1]
// Output: 1

nums = [2,2,1]

var singleNumber = function(nums) {
    let result = nums[0];
    for (let i = 1; i<nums.length;i++){
        result ^= nums[i];
    }
    return result;
};
