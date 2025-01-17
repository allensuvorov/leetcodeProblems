func minFlips(a int, b int, c int) int {
    all := (a | b) ^ c
    double := all & (a & b)

    return countBits(all) + countBits(double)
}

func countBits(num int) int {
    count := 0
    for num > 0 {
        count += num & 1
        num >>= 1
    }
    return count
}
