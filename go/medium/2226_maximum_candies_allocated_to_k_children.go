func maximumCandies(candies []int, k int64) int {
    l, r := 0, int(1e7)

    m := l + (r-l)/2
    for l < r {
        m = l + (r-l)/2 + 1

        if canSplit(candies, m, k) {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}

func canSplit(candies []int, guess int, k int64) bool {
    for _, v := range candies {
        k -= int64(v / guess)
    }
    return k <= 0
}
