func findSpecialInteger(arr []int) int {
    threshold := len(arr)*25/100
    fm := map[int]int{}
    for _, v := range arr {
        fm[v]++
        if fm[v] > threshold {
            return v
        }
    }
    return 0
}
