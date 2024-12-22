func longestConsecutive(nums []int) int {
    maxLenSeq := 0
    visited := map[int]bool{}
    exists := func(num int) bool {
        _, ok := visited[num]
        return ok
    }
    for _, v := range nums {
        visited[v] = false
    }
    for _, v := range nums {
        if visited[v] {
            continue
        }
        visited[v] = true
        curSeq := 1
        for up := v + 1; exists(up); up++{
            visited[up] = true
            curSeq++
        }
        for down := v - 1; exists(down); down-- {
            visited[down] = true
            curSeq++
        }
        maxLenSeq = max(maxLenSeq, curSeq)
    }
    return maxLenSeq
}
