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
    days--
    for _, v := range weights {
        if v > capacity {
            return false
        }
        
        if load + v > capacity {
            days--
            load = v
        } else {
            load += v
        }
    }
    return days >= 0
}

// l = 1
// r = 5000
// m = 2500
// Input: weights = [1,2,3,4,5,6,7,8,9,10], days = 5
