func maxDistance(colors []int) int {
    last := len(colors) - 1
    for i, j := 0, last; i <= j; i, j = i + 1, j - 1 {
        if colors[j] != colors[0] {
            return j-0
        }
        if colors[i] != colors[last] {
            return last - i
        }
    }
    return 0
}
