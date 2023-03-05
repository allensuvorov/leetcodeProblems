func maxDistance(colors []int) int {
    i := 0
    j := len(colors) - 1
    for i <= j {
        if colors[j] != colors[0] {
            return j-0
        }
        if colors[i] != colors[len(colors) - 1] {
            return len(colors) - 1 - i
        }
        i++
        j--
    }
    return 0
}
