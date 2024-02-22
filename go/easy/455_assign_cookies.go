func findContentChildren(g []int, s []int) int {
    slices.Sort(g)
    slices.Sort(s)
    ans := 0
    i, j := 0, 0
    for i < len(g) && j < len(s) {
        if g[i] > s[j] {
            j++
            continue
        }
        ans++
        j++
        i++
    }
    return ans
}
