func decrypt(code []int, k int) []int {
    // sliding window
    size := len(code)
    l, r := 0, 0
    res := make([]int, size)
    switch {
    case k == 0:
        return res
    case k > 0:
        l = 1
        r = k
    case k < 0:
        l = size + k // 4 - 
        r = size - 1
    }
    // fill the window
    window := 0
    for i := l; i <= r; i++ {
        window += code[i]
    }
    for i := range code {
        res[i] = window
        window -= code[l]
        l++
        r++
        if l >= size {
            l = 0
        }
        if r >= size {
            r = 0
        }
        window += code[r]
    }
    return res
}
