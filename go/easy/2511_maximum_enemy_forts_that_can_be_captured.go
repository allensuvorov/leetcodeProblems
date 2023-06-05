func captureForts(forts []int) int {
    ans := 0
    count := 0
    prevI := 0
    prev := 0
    for i, v := range forts {
        if v != 0 {
            if v == -prev {
                count = i - prevI - 1
                if ans < count {
                    ans = count
                }
            }
            prevI = i
            prev = v
        }
    }
    return ans
}
