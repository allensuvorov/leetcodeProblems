func heightChecker(heights []int) int {
    expected := make([]int, len(heights))
    copy(expected, heights)
    sort.Ints(expected)
    
    count := 0
    for i := range expected {
        if heights[i] != expected[i] {
            count++
        }
    }
    return count
}
