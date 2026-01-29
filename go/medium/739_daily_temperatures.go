// stack solution
func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    st := make([]int, 0) // stack of indices
    for i, v := range temperatures {
        for len(st) > 0 && temperatures[st[len(st) - 1]] < v {
            top := len(st) - 1
            res[st[top]] = i - st[top]
            st = st[:top] // pop
        }
        st = append(st, i)
    }
    return res
}

// bounded-temp allows for a more intuitive solution
// go right to left and track nextOccurrence
// index is temperature, value is index (day)
func dailyTemperatures(temperatures []int) []int {
    nextOccurrence := make([]int, 101)
    for i := range nextOccurrence {
        nextOccurrence[i] = math.MaxInt
    }
    n := len(temperatures)
    result := make([]int, n)
    // from right to left <-
    for i := n - 1; i >= 0; i-- {
        t := temperatures[i]
        nextOccurrence[t] = i
        // lookup nextOccurrence for a higher temperature
        minIndex := math.MaxInt
        for j := t + 1; j < len(nextOccurrence); j++ {
            minIndex = min(minIndex, nextOccurrence[j])
        }
        if minIndex != math.MaxInt {
            result[i] = minIndex - i
        }
    }
    return result
}
