func doesValidArrayExist(derived []int) bool {
    xorSum := 0
    for _, v := range derived {
        xorSum ^= v
    }
    return xorSum == 0 
}
