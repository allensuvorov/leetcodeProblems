func minEatingSpeed(piles []int, h int) int {
    l, r := 1, 1_000_000_000
    for l <= r { 
        m := l + (r-l)/2
        t := getEatingTime(piles, m) 
        if t <= h {
            r = m - 1 // slow down
        } else {
            l = m + 1 // speed up
        }
    }
    return l
}

func getEatingTime(piles []int, k int) int {
    t := 0
    for _, v := range piles {
        t += v / k
        if v % k > 0 {
            t++
        }
    }
    return t
}

// n * log n
    // max time <= h
    // h = 1,2,3,4,5,6,7,8,9,10

    // k - h               
    // 3 - 10
    // 4 - 8
    // 5 - 8
    // 6 - 6
