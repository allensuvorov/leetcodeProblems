// binary search and simple iterative count 

/*
Tastiness
              l m  r
1 2 3 4 5 6 7 8 9 10

Distanses
*                    *                             *
1 2     5     8            13                      21
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21

*/

func maximumTastiness(price []int, k int) int {
    slices.Sort(price)

    // range of possible tastiness
    l := 0
    r := (price[len(price) - 1] - price[0]) / (k - 1)

    for l < r {
        m := l + (r-l)/2 + 1 // a guess tastiness
        candyCount := countCandyAtTastiness(price, m)

        if candyCount >= k {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}

func countCandyAtTastiness(price []int, dist int) int {
    candyCount := 1
    lastPrice := price[0]
    for i := range price {
        if price[i] - lastPrice >= dist {
            candyCount++
            lastPrice = price[i]
        }
    }
    return candyCount
}
