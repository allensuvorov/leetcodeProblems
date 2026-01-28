func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    days := []int{}

    for i, v := range temperatures {
        for len(days) > 0 && temperatures[days[len(days)-1]] < v {
            top := len(days)-1
            res[days[top]] = i - days[top]
            days = days[:top]
        }
        days = append(days, i)
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
