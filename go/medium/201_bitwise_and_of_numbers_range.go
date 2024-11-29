func rangeBitwiseAnd(left int, right int) int {
    if left == right {
        return left
    }

    if left == 0 {
        return 0
    }

    log2Left := int(math.Log2(float64(left)))
    log2Right := int(math.Log2(float64(right)))

    if log2Left != log2Right {
        return 0
    }


    sum := int(math.Pow(2, float64(log2Left)))


    return sum + rangeBitwiseAnd(left - sum, right - sum)
}
