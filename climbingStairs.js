// https://leetcode.com/problems/climbing-stairs/
// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps. 
// In how many distinct ways can you climb to the top?
// Constraints: 1 <= n <= 45

/**
 * @param {number} n
 * @return {number}
 */

// bruteforce solution
var climbStairs = function(n) {
    let climbStairsRecursively = function(i,n) {
        if (i > n) {
            return 0;
        };
        if (i===n) {
            return 1;
        };
        return climbStairsRecursively(i+1,n) + climbStairsRecursively(i+2,n);
    };
    return climbStairsRecursively (0, n);
};
console.log(climbStairs(4)); 
// (3) 111, 21, 12 -> 3
// (4) 1111, 211, 121, 112, 22 -> 5
// (5) 11111, 2111, 1211, 1121, 1112, 221, 212, 122 -> 8
// (6) 111111, 21111*5, 2211, 222