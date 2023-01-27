func numberOfLines(widths []int, s string) []int {
    var sum, lineCount int
    for _, v := range s{
        w := widths[v - 'a']
        if sum + w > 100 {
            lineCount++
            sum = w
        } else {
            sum += w
        }
    }
    return []int{lineCount+1, sum}
}
