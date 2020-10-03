// https://leetcode.com/problems/plus-one/
// Given a non-empty array of digits representing a non-negative integer, increment one to the integer.
// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.

var plusOne = function(digits) {
    for (let i = digits.length -1; i>= 0; i--) {
        // console.log(digits[i]);
        if (digits[i]===9) {
            digits[i] = 0;
            // console.log(i,digits[i]);
            if (i === 0 && digits[i] === 0) {
                digits[i] = 1;
                digits.push(0);
                return digits;
            };
            } else { 
                digits[i] += 1;
                return digits;
            };
    };
};

console.log(plusOne([9,9,9]));
console.log(plusOne([1,2,3]));
console.log(plusOne([9]));
console.log(plusOne([0]));