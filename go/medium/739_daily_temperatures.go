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
