func numberOfAlternatingGroups(colors []int, k int) int {
    ans := 0
    n := len(colors)
    for l, r := 0, 1; r < n + k - 1; r++ {
        if colors[r%n] != colors[(r-1)%n] {
            if (r - l + 1) >= k {
                ans ++
            } 
        } else {
            l = r
        }
    }
    return ans
}
