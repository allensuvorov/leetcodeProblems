/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    // two pointers
    l := 1
    r := n
    // loop
    for l <= r {
        // gett the middle
        mid := (l + r)/2
        // move pointers
        switch guess(mid) {
        case 0: 
            return mid
        case 1: 
            l = mid + 1
        case -1: 
            r = mid - 1
        }
    }
    return -1
}
