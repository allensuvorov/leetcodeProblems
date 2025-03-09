func numberOfAlternatingGroups(colors []int) int {
    n := len(colors)
    ans := 0
    for i1 := range colors {
        i2 := i1 + 1
        if i2 >= n {
            i2 = i2 - n
        }

        i3 := i1 + 2
        if i3 >= n {
            i3 = i3 - n
        }

        if colors[i1] != colors[i2] && colors[i2] != colors[i3] {
            ans++
        }
    }
    return ans
}
