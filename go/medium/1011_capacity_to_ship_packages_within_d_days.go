func shipWithinDays(weights []int, days int) int {
    l, r := 1, len(weights) * 500

    for l < r {
        m := l + (r - l)/2 
        if canShip(weights, days, m) {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func canShip(weights []int, days, capacity int) bool {
    load := 0
    dayCount := 1
    for _, v := range weights {
        if v > capacity {
            return false
        }
        
        if load + v > capacity {
            load = 0
            dayCount++
        } 
        load += v
    }
    return days >= dayCount
}
