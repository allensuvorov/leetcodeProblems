// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Given a sorted array nums, remove the duplicates in-place such that each element 
// appears only once and returns the new length.

var removeDuplicates = function(nums) {
    let length = 0; // unique sequence index
    for (let i = 0; i <nums.length-1; i++) {
        let first = i; // grab index of first element
        while (nums[first] === nums[i+1]) { // untill we get to next unique element 
            i++; // one step forward
            if (i===nums.length-1) return length+1; // if hit the end index return current length
        };
        length++; // increase index for unique sequence
        nums[length] = nums[i+1]; // put next value to next unique
    };
    console.log(nums);
    return length+1;
};

let nums5 = [0,0,1,1,1,2,2,3,3,4,5];
console.log(removeDuplicates(nums5));