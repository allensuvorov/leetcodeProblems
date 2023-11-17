func findComplement(num int) int {
    res := 0
    for shift := 0; num != 0; shift++ {
        // extract
        bit := ^num & 1
        num >>= 1
        // collect
        res += bit << shift
    }
    return res
}
