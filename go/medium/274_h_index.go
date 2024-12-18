func hIndex(citations []int) int {
    const maxCitation = 1000
    counts := make([]int, maxCitation + 1)
    for _, v := range citations {
        counts[v]++
    }
    sum := 0
    c := maxCitation
    for c >= 0 && c > sum  {
        sum += counts[c]
        if c <= sum {
            return c
        }
        c--
    }
    return c
}
