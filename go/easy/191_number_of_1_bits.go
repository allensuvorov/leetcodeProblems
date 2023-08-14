func hammingWeight(num uint32) int {
    count := 0
    for i := 0; i < 32; i++ {
        count += int(num & 1)
        num >>= 1
    }
    return count
}
