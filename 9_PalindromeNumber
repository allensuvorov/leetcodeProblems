//Given an integer x, return true if x is palindrome integer.

//An integer is a palindrome when it reads the same backward as forward.

//For example, 121 is a palindrome while 123 is not.

/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    let y = x;
    let xRev = 0;
    if (y<0) return false;
    if (y>=0 && y<10) return true;
    while (y>=10) {
        xRev += y%10; //
        xRev *= 10; //
        y = Math.floor(y/10); //
    }
    xRev += y%10;
    return (x === xRev);

};
