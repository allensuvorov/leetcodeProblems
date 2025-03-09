func numberOfAlternatingGroups(colors []int) int {
    n := len(colors)
    ans := 0
    for i := range colors {
        if colors[i] != colors[(i+1)%n] && colors[(i+1)%n] != colors[(i+2)%n] {
            ans++
        }
    }
    return ans
}
