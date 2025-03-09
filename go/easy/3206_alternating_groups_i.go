func numberOfAlternatingGroups(colors []int) int {
    ans := 0
    double := append(colors, colors...)
    for i := range colors {
        if double[i] != double[i+1] && double[i+1] != double[i+2] {
            ans++
        }
    }
    return ans
}
