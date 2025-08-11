func reorderedPowerOf2(n int) bool {
    need := count(n)
    for i := range 32 {
        if count(1 << i) == need {
            return true
        }
    }
    return false
}

func count(num int) [10]int {
    counts := [10]int{}
    for num > 0 {
        counts[num % 10]++
        num = num / 10
    }
    return counts
}
