func binaryGap(n int) int {
    maxGap := 0
    curGap := 0
    inGap := false
    for n > 0 {
        d := n & 1
        if inGap {
            if d == 0 {
                curGap++
            }
            if d == 1 {
                maxGap = max(maxGap, curGap + 1)
                curGap = 0
            }
        } else {
            if d == 1 {
                inGap = true
            }
        }
        n >>= 1
    }

    return maxGap
}
