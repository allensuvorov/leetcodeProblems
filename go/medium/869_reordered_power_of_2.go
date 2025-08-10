func reorderedPowerOf2(n int) bool {
    // try each order Big O((Log(10)N)!)
    digitCounts := make(map[int]int)
    for x := n; x > 0; x = x / 10 {
        digitCounts = append(digits, x % 10)
    }

    return dfs(digitCounts, num)
}

func dfs(digitCounts map[int]int, num int) bool {
    if len(digitCounts) == 0 {
        return isPowerOf2(num)
    }

    for i := range len(digits) {
        if digits[i] != 0 {
            used := make([]bool, len(digits))
            num := digits[i]
            used[i] = true
        }
    }

    return false
}

func isPowerOf2(n int)bool {
    return n > 0 && (n & (n-1) == 0)
}

1    2    3
2 3  1 3
3 2  3 1
