func minEatingSpeed(piles []int, h int) int {
    l := 1
    r := 1_000_000_000
    for l <= r {
        m := l + (r - l)/2
        time := eatTime(piles, m)
        if time <= h {
            r = m-1
        } else {
            l = m+1
        }
    }
    return l
}


func eatTime(piles []int, k int) int {
    hours := 0
    for _, v := range piles {
        if v % k != 0 {
            hours++
        }
        hours += v / k
    }
    return hours
}
