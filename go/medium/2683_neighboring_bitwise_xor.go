func doesValidArrayExist(derived []int) bool {
    xorSum := 0
    for _, v := range derived {
        xorSum ^= v
    }
    return xorSum == 0 
}

// reverse
func doesValidArrayExist(derived []int) bool {
    first := 1
    curr := 1
    for _, v := range derived {
        if v == 1 {
            curr = -curr
        }
    }
    return curr == f
}
