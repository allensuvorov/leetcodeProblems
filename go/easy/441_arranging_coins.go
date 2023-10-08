// Solution 1 - math - (k * (k + 1)) / 2 <= n
func arrangeCoins(n int) int {
    return int(math.Sqrt(2.0 * float64(n) + 0.25) - 0.5)
}

// Solution 2 - iteration
func arrangeCoins(n int) int {
    sum := 0
    inc := 1
    i := 0
    for ; sum < n; i++  {
        sum += inc
        inc ++
        if sum > n {
            return i
        }
    }
    return i
}

// Solution 3 - Binary search
func arrangeCoins(n int) int {
    l := 1
    r := n
    for l < r {
        m := l + (r - l + 1)/2
        if countCoins(m) <= n {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}

func countCoins(rows int) (coins int) {
    coins = (1 + rows) * rows / 2
    return coins
}
