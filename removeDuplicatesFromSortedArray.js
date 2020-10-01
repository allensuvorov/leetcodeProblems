// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Given a sorted array nums, remove the duplicates in-place such that each element 
// appears only once and returns the new length.

var removeDuplicates = function(nums) {
    let length = 0; // unique sequence index
    
    for (let i = 0; i <nums.length; i++) {
        let first = i; // grab index of first element
        while (nums[first] === nums[i+1]) { // untill we get to next unique element 
            console.log('here in while');
            i++; // one step forward
        };
        console.log(`i+1 = ${i+1}, nums[i+1] = ${nums[i+1]}`);
        if (i !== first) { // if found dubplicate on that i step
            length++; // increase index for unique sequence
            nums[length] = nums[i+1]; // put next value to next unique
            
            console.log(length,i, nums[i], nums);
        };
        

    };

    return length+1;
};

let nums5 = [0,0,1,1,1,2,2,3,3,4];

console.log(removeDuplicates(nums5));