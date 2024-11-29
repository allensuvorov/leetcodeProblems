func rangeBitwiseAnd(left int, right int) int {
    if left == right {
        return left
    }
    left >>= 1
    right >>= 1
    return rangeBitwiseAnd(left, right) << 1
}
