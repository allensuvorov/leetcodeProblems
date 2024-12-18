// count sort
func hIndex(citations []int) int {
    counts := make([]int, 1000 + 1)
    for _, v := range citations {
        counts[v]++
    }
    sum := 0
    for i := len(counts) - 1; i >= 0; i-- {
        sum += counts[i]
        if sum >= i {
            return i
        }
    }
    return 0
}

// general sort
func hIndex(citations []int) int {
    n := len(citations)
    slices.Sort(citations)
    i := n - 1
    ans := 0
    for i >= 0 && citations[i] >= n - i {
        ans++
        i--
    }
    return ans
}
