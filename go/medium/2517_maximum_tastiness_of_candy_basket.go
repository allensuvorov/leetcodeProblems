func maximumTastiness(price []int, k int) int {
    slices.Sort(price)

    l := 0
    r := (price[len(price) - 1] - price[0]) / (k - 1)

    for l < r {
        m := l + (r-l)/2 + 1
        candyCount := fillsBasket(price, m)

        if candyCount >= k {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}

func fillsBasket(price []int, dist int) int {
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
