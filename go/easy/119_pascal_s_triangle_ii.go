func getRow(rowIndex int) []int {
    prev := []int{}
    for i := 0; i <= rowIndex; i++ {
        row := make([]int, i + 1)
        row[0] = 1
        row[i] = 1
        for j := 1; j < i; j++ {
            row[j] = prev[j - 1] + prev[j]
        }
        prev = row
    }
    return prev
}
