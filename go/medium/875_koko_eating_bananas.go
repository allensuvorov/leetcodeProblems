func minEatingSpeed(piles []int, h int) int {
    l := len(piles)
    r := 1_000_000_000
    for l < r {
        m := l + (r - l)/2 + 1
        time := eatTime(piles, m)
        if time <= h {
            r = m-1
        } else {
            l = m
        }
    }
    return l
}


func eatTime(piles []int, k int) int {
    count := 0
    for _, v := range piles {
        count += v / k + 1
    }
    return count
}
