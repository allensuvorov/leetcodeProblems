/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 1, n
    for l <= r {
        m := l + (r - l)/2
        switch guess(m) {
        case 0:
            return m
        case -1:
            r = m - 1
        case 1:
            l = m + 1
        }
    }
    return -1
}

/*
          p
1 2 3 4 5 6 7 8 9 10 
          m
            r
          l
*/
